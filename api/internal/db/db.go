package db

import (
	"simple-setup/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

	if err != nil {
		return nil
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Group{}, &models.User{})
}
