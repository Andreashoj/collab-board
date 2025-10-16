package handlers

import "github.com/google/uuid"

type CreateLogRequest struct {
	BoardID uuid.UUID `json:"board_id"`
	UserID  uuid.UUID `json:"user_id"`
	Change  string    `json:"change"`
}

type BoardLogResponse struct {
	ID      uuid.UUID `json:"id"`
	BoardID uuid.UUID `json:"board_id"`
	UserID  uuid.UUID `json:"user_id"`
	Change  string    `json:"change"`
}
