package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"simple-setup/internal/models"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

type contextKey string

const UserContextKey contextKey = "user"

type AuthService struct {
	firebaseAuth *auth.Client
	db           *gorm.DB
}

func NewAuthService(database *gorm.DB) (*AuthService, error) {
	serviceAccountJSON := os.Getenv("FIREBASE_SERVICE_ACCOUNT_JSON")
	if serviceAccountJSON == "" {
		return nil, errors.New("FIREBASE_SERVICE_ACCOUNT_JSON environment variable not set")
	}

	opt := option.WithCredentialsJSON([]byte(serviceAccountJSON))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	firebaseAuth, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	log.Println("Firebase Admin SDK initialized")

	return &AuthService{
		firebaseAuth: firebaseAuth,
		db:           database,
	}, nil
}

// Middleware returns an http middleware that verifies Firebase ID tokens
func (a *AuthService) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		idToken := parts[1]

		// Verify the token with request context
		token, err := a.firebaseAuth.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			log.Printf("Error verifying ID token: %v", err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Ensure user exists in database (auto-create if not)
		a.ensureUserExists(token)

		// Add user info to context
		ctx := context.WithValue(r.Context(), UserContextKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserFromContext retrieves the Firebase user from request context
func GetUserFromContext(ctx context.Context) (*auth.Token, error) {
	user, ok := ctx.Value(UserContextKey).(*auth.Token)

	if !ok {
		return nil, errors.New("Couldn't get user from token")
	}

	return user, nil
}

// ensureUserExists checks if user exists in database, creates if not
func (a *AuthService) ensureUserExists(token *auth.Token) {
	// Check if user exists by Firebase UID
	var existingUser models.User
	result := a.db.Where("firebase_uid = ?", token.UID).First(&existingUser)

	if result.Error == gorm.ErrRecordNotFound {
		// User doesn't exist, create them
		email := ""
		if emailVal, ok := token.Claims["email"].(string); ok {
			email = emailVal
		}

		newUser := models.User{
			FirebaseUID:  token.UID,
			Email:        email,
			PasswordHash: "",
		}

		if err := a.db.Create(&newUser).Error; err != nil {
			log.Printf("Failed to create user in database: %v", err)
		} else {
			log.Printf("Created new user in database: %s (%s)", email, token.UID)
		}
	}
}
