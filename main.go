package main

import (
	"log"
	"net/http"
	"os"

	"model/config"
	"model/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		// Don't fatal here for serverless - let it continue
		return
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
	router.Run("0.0.0.0:10000")
}
