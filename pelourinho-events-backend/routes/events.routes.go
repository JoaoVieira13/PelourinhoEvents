package routes

import (
	"pe/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetRoutes(r *mux.Router, db *gorm.DB) {
	eventService := services.NewEventService(db)
	r.HandleFunc("/events", eventService.GetEventsHandler()).Methods("GET")
	r.HandleFunc("/events", services.CreateEvent(db)).Methods("POST")
	r.HandleFunc("/events/{id}", services.DeleteEvent(db)).Methods("DELETE")
	r.HandleFunc("/events/{id}", services.UpdateEvent(db)).Methods("PUT")
}
