package routes

import (
	"pe/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/events", services.GetEvents(db)).Methods("GET")
	r.HandleFunc("/events", services.CreateEvent(db)).Methods("POST")
}
