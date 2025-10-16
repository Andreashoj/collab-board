package repositories

import (
	"simple-setup/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// Repository
func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{db: database}
}

// Crud
func (r *UserRepository) CreateUser(user *models.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserById(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *UserRepository) DeleteUser(id uuid.UUID) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) EditUser(u *models.User) error {
	if err := r.db.Save(&u).Error; err != nil {
		return err
	}

	return nil
}
