package handlers

import "github.com/google/uuid"

type AddMemberRequest struct {
	UserID  uuid.UUID `json:"user_id"`
	BoardID uuid.UUID `json:"board_id"`
	Role    string    `json:"role"` // "owner", "viewer", "editor"
}

type UpdateMemberRoleRequest struct {
	Role string `json:"role"`
}

type BoardMemberResponse struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	BoardID uuid.UUID `json:"board_id"`
	Role    string    `json:"role"`
}
