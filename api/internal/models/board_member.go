package models

import (
	"time"

	"github.com/google/uuid"
)

type BoardMember struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	BoardID   uuid.UUID `gorm:"type:uuid;not null;index" json:"board_id"`
	Role      string    `gorm:"not null;default:'viewer'" json:"role"` // "owner", "viewer", "editor"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Relations
	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Board Board `gorm:"foreignKey:BoardID" json:"board,omitempty"`
}
