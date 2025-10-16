package services

import (
	"simple-setup/internal/models"
	"simple-setup/internal/repositories"

	"github.com/google/uuid"
)

type BoardService struct {
	repository *repositories.BoardRepository
}

func NewBoardService(repo *repositories.BoardRepository) *BoardService {
	return &BoardService{repository: repo}
}

func (s *BoardService) CreateBoard(board *models.Board) error {
	return s.repository.CreateBoard(board)
}

func (s *BoardService) GetBoard(id uuid.UUID) (*models.Board, error) {
	return s.repository.GetBoardByID(id)
}

func (s *BoardService) GetAllBoards() (*[]models.Board, error) {
	return s.repository.GetAllBoards()
}

func (s *BoardService) GetBoardsByUser(userID uuid.UUID) (*[]models.Board, error) {
	return s.repository.GetBoardsByUserID(userID)
}

func (s *BoardService) UpdateBoard(board *models.Board) error {
	return s.repository.UpdateBoard(board)
}

func (s *BoardService) DeleteBoard(id uuid.UUID) error {
	return s.repository.DeleteBoard(id)
}
