package models

import "time"

type Event struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Title         string    `json:"title"`
	SubTitle      string    `json:"subTitle"`
	EventDate     time.Time `json:"eventDate"`
	EventLocation string    `json:"eventLocation"`
	Rate          int       `json:"rate"`
}
