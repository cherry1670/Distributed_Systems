package models

import (
	"github.com/google/uuid"
)

type Task_Queue struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Assigned string    `gorm:"type:ENUM('pending', 'in_progress', 'completed', 'failed');not null"`
	Priority int       `gorm:"not null"`
	TaskID   uuid.UUID `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Task     Task      `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
