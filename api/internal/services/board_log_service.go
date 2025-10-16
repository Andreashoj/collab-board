package services

import (
	"simple-setup/internal/models"
	"simple-setup/internal/repositories"

	"github.com/google/uuid"
)

type BoardLogService struct {
	repository *repositories.BoardLogRepository
}

func NewBoardLogService(repo *repositories.BoardLogRepository) *BoardLogService {
	return &BoardLogService{repository: repo}
}

func (s *BoardLogService) CreateLog(log *models.BoardLog) error {
	return s.repository.CreateLog(log)
}

func (s *BoardLogService) GetLogsByBoard(boardID uuid.UUID) (*[]models.BoardLog, error) {
	return s.repository.GetLogsByBoardID(boardID)
}

func (s *BoardLogService) GetLog(id uuid.UUID) (*models.BoardLog, error) {
	return s.repository.GetLogByID(id)
}

func (s *BoardLogService) DeleteLog(id uuid.UUID) error {
	return s.repository.DeleteLog(id)
}
