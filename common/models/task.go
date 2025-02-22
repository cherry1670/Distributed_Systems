package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Type      string    `gorm:"type:ENUM('image_processing', 'file_operations');not null"`
	Status    string    `gorm:"type:ENUM('pending', 'in_progress', 'completed', 'failed');not null"`
	Priority  int       `gorm:"not null"`
	InputData string    `gorm:"type:jsonb;not null"` // Use `jsonb` for PostgreSQL
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (Task) TableName() string {
	return "task"
}
