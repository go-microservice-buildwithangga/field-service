package dto

import (
	"time"

	"github.com/google/uuid"

	"field-service/constants"
)

type FieldScheduleRequest struct {
	FieldID uint     `json:"fieldID" validate:"required"`
	Date    string   `json:"date" validate:"required"`
	TimeID  []string `json:"timeIDs" validate:"required"`
}

type GenerateFieldScheduleFromOneMonthRequest struct {
	FieldID uint `json:"fieldID" validate:"required"`
}

type UpdateFieldScheduleRequest struct {
	Date   string   `json:"date" validate:"required"`
	TimeID []string `json:"timeIDs" validate:"required"`
}

type UpdateStatusScheduleRquest struct {
	FieldScheduleIDs []string `json:"fieldScheduleIDs" validate:"required"`
}

type FieldScheduleResponse struct {
	UUID         uuid.UUID                         `json:"uuid"`
	FieldName    string                            `json:"fieldName"`
	PricePerHour int                               `json:"pricePerHour"`
	Date         string                            `json:"date"`
	Status       constants.FieldScheduleStatusName `json:"status"`
	Time         string                            `json:"time"`
	CreatedAt    *time.Time                        `json:"createdAt"`
	UpdatedAt    *time.Time                        `json:"updatedAt"`
}

type FieldScheduleBookingResponse struct {
	UUID         uuid.UUID                         `json:"uuid"`
	PricePerHour string                            `json:"pricePerHour"`
	Date         string                            `json:"date"`
	Status       constants.FieldScheduleStatusName `json:"status"`
	Time         string                            `json:"time"`
}

type FieldScheduleRequestParam struct {
	Page       int     `form:"page" validate:"required"`
	Limit      int     `form:"limit" validate:"required"`
	SortColumn *string `form:"sortColumn"`
	SortOrder  *string `form:"sortOrder"`
}

type FieldScheduleByFieldIDAndDateRequestParam struct {
	FieldID uint   `form:"fieldID" validate:"required"`
	Date    string `form:"date" validate:"required"`
}
