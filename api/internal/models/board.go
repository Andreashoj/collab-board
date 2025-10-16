package models

import (
	"time"

	"github.com/google/uuid"
)

type Board struct {
	ID        uuid.UUID     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string        `gorm:"not null" json:"name"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Members   []BoardMember `gorm:"foreignKey:BoardID" json:"members,omitempty"`
	Logs      []BoardLog    `gorm:"foreignKey:BoardID" json:"logs,omitempty"`
}
