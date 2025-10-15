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

	// Firebase
	if err := middlewares.InitFirebase(); err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	// Db
	database := db.Init()

	// Redis
	redisclient.NewClient()

	// Repos
	userRepository := repositories.NewUserRepository(database)

	// Services
	userService := services.NewUserService(userRepository)

	// Handlers
	userHandler := handlers.NewUserHandler(userService)

	router := chi.NewRouter()

	// Middlewares
	middlewares.DefineMiddleware(router)

	// Routes
	userHandler.RegisterRoutes(router)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
