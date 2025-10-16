package repositories

import (
	"simple-setup/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardMemberRepository struct {
	db *gorm.DB
}

func NewBoardMemberRepository(database *gorm.DB) *BoardMemberRepository {
	return &BoardMemberRepository{db: database}
}

func (r *BoardMemberRepository) AddMember(member *models.BoardMember) error {
	return r.db.Create(member).Error
}

func (r *BoardMemberRepository) GetMembersByBoardID(boardID uuid.UUID) (*[]models.BoardMember, error) {
	var members []models.BoardMember
	if err := r.db.Preload("User").Where("board_id = ?", boardID).Find(&members).Error; err != nil {
		return nil, err
	}
	return &members, nil
}

func (r *BoardMemberRepository) GetMemberByID(id uuid.UUID) (*models.BoardMember, error) {
	var member models.BoardMember
	if err := r.db.Preload("User").Preload("Board").First(&member, id).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *BoardMemberRepository) UpdateMemberRole(id uuid.UUID, role string) error {
	return r.db.Model(&models.BoardMember{}).Where("id = ?", id).Update("role", role).Error
}

func (r *BoardMemberRepository) RemoveMember(id uuid.UUID) error {
	return r.db.Delete(&models.BoardMember{}, id).Error
}

func (r *BoardMemberRepository) GetMemberByUserAndBoard(userID, boardID uuid.UUID) (*models.BoardMember, error) {
	var member models.BoardMember
	if err := r.db.Where("user_id = ? AND board_id = ?", userID, boardID).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}