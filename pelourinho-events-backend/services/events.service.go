package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"pe/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventService struct {
	db *gorm.DB
}

func NewEventService(db *gorm.DB) *EventService {
	return &EventService{
		db: db,
	}
}

// func GetEvents(db *gorm.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var events []models.Event
// 		db.Find(&events)
// 		json.NewEncoder(w).Encode(events)
// 	}
// }

func GetEvents(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var event []models.Event

	err := db.Find(&event).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateEvent(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event models.Event
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		event.ID = uuid.New()

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
