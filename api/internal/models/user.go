package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	FirebaseUID  string        `gorm:"unique;not null" json:"firebase_uid"`
	Email        string        `gorm:"unique;not null" json:"email"`
	PasswordHash string        `json:"-"` // Never expose in JSON
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	Boards       []BoardMember `gorm:"foreignKey:UserID" json:"boards,omitempty"`
}
