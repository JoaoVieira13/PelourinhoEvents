package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Title         string    `json:"title"`
	SubTitle      string    `json:"subTitle"`
	EventDate     time.Time `json:"eventDate"`
	EventLocation string    `json:"eventLocation"`
	Rate          int       `json:"rate"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	ModifiedAt    time.Time `json:"modifiedAt" gorm:"autoUpdateTime"`
}
