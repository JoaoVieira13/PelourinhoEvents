package services

import (
	"encoding/json"
	"net/http"
	"pe/models"

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
		db.Create(&event)
		json.NewEncoder(w).Encode(event)
	}
}
