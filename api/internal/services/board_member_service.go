package services

import (
	"simple-setup/internal/models"
	"simple-setup/internal/repositories"

	"github.com/google/uuid"
)

type BoardMemberService struct {
	repository *repositories.BoardMemberRepository
}

func NewBoardMemberService(repo *repositories.BoardMemberRepository) *BoardMemberService {
	return &BoardMemberService{repository: repo}
}

func (s *BoardMemberService) AddMember(member *models.BoardMember) error {
	return s.repository.AddMember(member)
}

func (s *BoardMemberService) GetMembersByBoard(boardID uuid.UUID) (*[]models.BoardMember, error) {
	return s.repository.GetMembersByBoardID(boardID)
}

func (s *BoardMemberService) GetMember(id uuid.UUID) (*models.BoardMember, error) {
	return s.repository.GetMemberByID(id)
}

func (s *BoardMemberService) UpdateMemberRole(id uuid.UUID, role string) error {
	return s.repository.UpdateMemberRole(id, role)
}

func (s *BoardMemberService) RemoveMember(id uuid.UUID) error {
	return s.repository.RemoveMember(id)
}

func (s *BoardMemberService) GetMemberByUserAndBoard(userID, boardID uuid.UUID) (*models.BoardMember, error) {
	return s.repository.GetMemberByUserAndBoard(userID, boardID)
}
