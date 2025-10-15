package services

import (
	"simple-setup/internal/models"
	"simple-setup/internal/repositories"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repository: repo}
}

func (s *UserService) CreateUser(u *models.User) error {
	return s.repository.CreateUser(u)
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	user, err := s.repository.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.repository.EditUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repository.DeleteUser(id)
}
