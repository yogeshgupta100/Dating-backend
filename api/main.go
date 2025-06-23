package main

import (
	"log"
	"net/http"
	"os"

	"model/config"
	"model/routes"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize router
	router = routes.SetupRouter(db)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Create a new Gin context
	ctx := gin.CreateTestContextOnly(w, router)
	ctx.Request = r

	// Handle the request
	router.HandleContext(ctx)
}

func main() {
	// This is for local development
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize router
	router := routes.SetupRouter(db)

	// Start server
	log.Printf("Server starting on port %s", port)
	router.Run(":" + port)
}
