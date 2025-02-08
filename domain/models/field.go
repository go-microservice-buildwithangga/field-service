package models

import (
	"time"

	"github.com/google/uuid"
)

type Field struct {
	ID             uint      `gorm:"primaryKey,autoIncrement"`
	UUID           uuid.UUID `gorm:"type:uuid;default:not null;unique"`
	Code           string    `gorm:"type:varchar(15);not null"`
	Name           string    `gorm:"type:varchar(50);not null"`
	PricePerHour   int       `gorm:"type:int;not null"`
	Images         string    `gorm:"type:text[];not null"`
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	FieldSchedules []FieldSchedule `gorm:"foreignKey:FieldID; references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
