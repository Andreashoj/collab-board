package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-setup/internal/db"
	"simple-setup/internal/handlers"
	"simple-setup/internal/middlewares"
	"simple-setup/internal/redis"
	"simple-setup/internal/repositories"
	"simple-setup/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Db
	database := db.Init()

	// AuthService & FireAuth
	authService, err := middlewares.NewAuthService(database)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	// Redis
	redisclient.NewClient()

	// Repos
	userRepository := repositories.NewUserRepository(database)
	boardRepository := repositories.NewBoardRepository(database)
	boardMemberRepository := repositories.NewBoardMemberRepository(database)
	boardLogRepository := repositories.NewBoardLogRepository(database)

	// Services
	userService := services.NewUserService(userRepository)
	boardService := services.NewBoardService(boardRepository)
	boardMemberService := services.NewBoardMemberService(boardMemberRepository)
	boardLogService := services.NewBoardLogService(boardLogRepository)

	// Handlers
	userHandler := handlers.NewUserHandler(userService, authService)
	boardHandler := handlers.NewBoardHandler(boardService, authService)
	boardMemberHandler := handlers.NewBoardMemberHandler(boardMemberService, authService)
	boardLogHandler := handlers.NewBoardLogHandler(boardLogService, authService)

	router := chi.NewRouter()

	// Middlewares
	middlewares.DefineMiddleware(router)

	// Routes
	userHandler.RegisterRoutes(router)
	boardHandler.RegisterRoutes(router)
	boardMemberHandler.RegisterRoutes(router)
	boardLogHandler.RegisterRoutes(router)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
