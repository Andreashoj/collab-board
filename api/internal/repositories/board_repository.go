package repositories

import (
	"simple-setup/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardRepository struct {
	db *gorm.DB
}

func NewBoardRepository(database *gorm.DB) *BoardRepository {
	return &BoardRepository{db: database}
}

func (r *BoardRepository) CreateBoard(board *models.Board) error {
	return r.db.Create(board).Error
}

func (r *BoardRepository) GetBoardByID(id uuid.UUID) (*models.Board, error) {
	var board models.Board
	if err := r.db.Preload("Members").Preload("Logs").First(&board, id).Error; err != nil {
		return nil, err
	}
	return &board, nil
}

func (r *BoardRepository) GetAllBoards() (*[]models.Board, error) {
	var boards []models.Board
	if err := r.db.Preload("Members").Find(&boards).Error; err != nil {
		return nil, err
	}
	return &boards, nil
}

func (r *BoardRepository) GetBoardsByUserID(userID uuid.UUID) (*[]models.Board, error) {
	var boards []models.Board
	if err := r.db.
		Joins("JOIN board_members ON board_members.board_id = boards.id").
		Where("board_members.user_id = ?", userID).
		Preload("Members").
		Find(&boards).Error; err != nil {
		return nil, err
	}
	return &boards, nil
}

func (r *BoardRepository) UpdateBoard(board *models.Board) error {
	return r.db.Save(board).Error
}

func (r *BoardRepository) DeleteBoard(id uuid.UUID) error {
	return r.db.Delete(&models.Board{}, id).Error
}
