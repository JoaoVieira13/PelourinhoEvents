package main

import (
	"log"
	"net/http"

	"pe/infra"
	"pe/routes"
	"pe/services"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db := infra.CreateConnection()

	services.NewEventService(db)
	// call controller

	// Init router
	r := mux.NewRouter()
	routes.SetRoutes(r, db)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(r)))
}
