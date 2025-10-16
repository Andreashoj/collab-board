package repositories

import (
	"simple-setup/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardLogRepository struct {
	db *gorm.DB
}

func NewBoardLogRepository(database *gorm.DB) *BoardLogRepository {
	return &BoardLogRepository{db: database}
}

func (r *BoardLogRepository) CreateLog(log *models.BoardLog) error {
	return r.db.Create(log).Error
}

func (r *BoardLogRepository) GetLogsByBoardID(boardID uuid.UUID) (*[]models.BoardLog, error) {
	var logs []models.BoardLog
	if err := r.db.Preload("User").Where("board_id = ?", boardID).Order("created_at desc").Find(&logs).Error; err != nil {
		return nil, err
	}
	return &logs, nil
}

func (r *BoardLogRepository) GetLogByID(id uuid.UUID) (*models.BoardLog, error) {
	var log models.BoardLog
	if err := r.db.Preload("User").Preload("Board").First(&log, id).Error; err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *BoardLogRepository) DeleteLog(id uuid.UUID) error {
	return r.db.Delete(&models.BoardLog{}, id).Error
}
