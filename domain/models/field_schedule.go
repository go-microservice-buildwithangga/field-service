package models

import (
	"time"

	"github.com/google/uuid"

	"field-service/constants"
)

type FieldSchedule struct {
	ID        uint                          `gorm:"primaryKey,autoIncrement"`
	UUID      uuid.UUID                     `gorm:"type:uuid;default:not null;unique"`
	FieldID   uint                          `gorm:"type:int;not null"`
	TimeID    uint                          `gorm:"type:int;not null"`
	Date      time.Time                     `gorm:"type:date;not null"`
	Status    constants.FieldScheduleStatus `gorm:"type:int;not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Field     Field `gorm:"foreignKey:FieldID; references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Time      Time  `gorm:"foreignKey:timeID; references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
