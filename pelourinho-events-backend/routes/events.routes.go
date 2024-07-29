package routes

import (
	"net/http"
	"pe/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		services.GetEvents(db, w, r)
	}).Methods("GET")
	r.HandleFunc("/events", services.CreateEvent(db)).Methods("POST")
	r.HandleFunc("/events/{id}", services.DeleteEvent(db)).Methods("DELETE")
	r.HandleFunc("/events/{id}", services.UpdateEvent(db)).Methods("PUT")
}
