package main

// Import necessary packages
import (
	"log"               // For logging messages to the console
	"net/http"          // For HTTP server functionality
	"server/handlers"   // Custom package for handling HTTP requests
	"server/services"   // Custom package for service-related logic
	"github.com/gorilla/mux"   // Third-party package for advanced routing
)

func main() {
	// Load store data
	services.InitStoreData("store.json")

	// Create a new router instance using Gorilla Mux
	// This router will handle HTTP request routing for the application
	router := mux.NewRouter()

	// Register endpoints
	router.HandleFunc("/api/submit/", handlers.SubmitJob).Methods("POST")   // This endpoint listens for POST requests at "/api/submit/"
	router.HandleFunc("/api/status", handlers.GetJobStatus).Methods("GET")  // This endpoint listens for GET requests at "/api/status"

	// Start the HTTP server
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
