package models

import (
	"time"

	"github.com/google/uuid"
)

type BoardLog struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	BoardID   uuid.UUID `gorm:"type:uuid;not null;index" json:"board_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Change    string    `gorm:"type:text;not null" json:"change"`
	CreatedAt time.Time `json:"created_at"`
	
	// Relations
	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Board Board `gorm:"foreignKey:BoardID" json:"board,omitempty"`
}
