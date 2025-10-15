package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type contextKey string

const UserContextKey contextKey = "user"

var firebaseAuth *auth.Client

// InitFirebase initializes Firebase Admin SDK from environment variable
func InitFirebase() error {
	serviceAccountJSON := os.Getenv("FIREBASE_SERVICE_ACCOUNT_JSON")
	if serviceAccountJSON == "" {
		return errors.New("FIREBASE_SERVICE_ACCOUNT_JSON environment variable not set")
	}

	opt := option.WithCredentialsJSON([]byte(serviceAccountJSON))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	firebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		return err
	}

	log.Println("Firebase Admin SDK initialized")
	return nil
}

// AuthMiddleware verifies Firebase ID tokens
func AuthMiddleware(next http.Handler) http.Handler {
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
		token, err := firebaseAuth.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			log.Printf("Error verifying ID token: %v", err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

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
