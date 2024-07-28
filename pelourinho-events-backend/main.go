package main

import (
	"log"
	"net/http"
	"os"

	"pe/models"
	"pe/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	// Get environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbPort == "" {
		log.Fatal("Database environment variables are not set properly")
	}
	// Define MySQL database connection details
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Init database
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	// Migrate schema
	db.AutoMigrate(&models.Event{})

	// Init router
	r := mux.NewRouter()
	routes.SetRoutes(r, db)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(r)))
}
