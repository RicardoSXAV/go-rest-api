package main

import (
	"go-rest-api/internal/database"
	"go-rest-api/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database connection
	db, err := database.NewClient()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Check and log initial migration status
	version, dirty, err := database.GetMigrationVersion(db)
	if err != nil {
		log.Printf("Migration status check failed: %v", err)
	} else {
		log.Printf("Migration status: version %d, dirty: %v", version, dirty)
	}

	// Initialize the router
	router := mux.NewRouter()

	// Initialize handlers
	orderHandler := handlers.NewOrderHandler(db)

	// Define the routes
	router.HandleFunc("/api/orders", orderHandler.CreateOrder).Methods("POST")
	router.HandleFunc("/api/orders", orderHandler.GetOrders).Methods("GET")

	// Start the server
	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Failed to start server", err)
	}
}
