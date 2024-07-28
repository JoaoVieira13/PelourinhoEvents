package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"pe/models"
	"time"

	"gorm.io/gorm"
)

func GetEvents(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var events []models.Event
		db.Find(&events)
		json.NewEncoder(w).Encode(events)
	}
}

func CreateEvent(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event models.Event
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		loc := time.Now()
		event.CreatedAt = loc

		db.Create(&event)
		json.NewEncoder(w).Encode(event)
	}
}

func DeleteEvent(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		var event models.Event
		db.First(&event, id)
		db.Delete(&event)
		json.NewEncoder(w).Encode(event)
	}
}

func UpdateEvent(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id") // Get the ID from the URL parameters
		var event models.Event

		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Find the existing event in the database by ID
		var existingEvent models.Event
		if err := db.First(&existingEvent, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, "Event not found", http.StatusNotFound)
			} else {
				http.Error(w, "Error finding event", http.StatusInternalServerError)
			}
			return
		}

		existingEvent.Title = event.Title
		existingEvent.SubTitle = event.SubTitle
		existingEvent.EventDate = event.EventDate
		existingEvent.EventLocation = event.EventLocation
		existingEvent.Rate = event.Rate
		existingEvent.ModifiedAt = time.Now()

		db.Save(&existingEvent)
	}
}
