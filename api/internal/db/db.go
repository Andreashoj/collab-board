package db

import (
	"fmt"
	"log"
	"os"
	"simple-setup/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// Get database connection details from environment variables
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "collabboard")
	password := getEnv("DB_PASSWORD", "collabboard_dev")
	dbname := getEnv("DB_NAME", "collabboard")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil
	}

	// Enable uuid-ossp extension
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	migrate(db)

	log.Println("Database connected successfully")
	return db
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Board{},
		&models.BoardMember{},
		&models.BoardLog{},
	)
}
