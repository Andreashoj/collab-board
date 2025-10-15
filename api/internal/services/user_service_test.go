package services

import (
	"simple-setup/internal/models"
	"simple-setup/internal/repositories"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Group{}); err != nil {
		t.Fatalf("failed to migrate models: %v", err)
	}

	return db
}

func TestCreateAndFindUser(t *testing.T) {
	db := setupTestDB(t)
	userRepo := repositories.NewUserRepository(db)
	userService := NewUserService(userRepo)

	// Arrange
	user := &models.User{
		Name:  "Andreas",
		Email: "høj@north.com",
	}

	// Act
	_, err := userService.CreateUser(user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
