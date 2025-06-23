package main

import (
	"log"
	"model/config"
	"model/routes"
)

func main() {
	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize router
	router := routes.SetupRouter(db)

	// Start server
	router.Run(":8080")
}
