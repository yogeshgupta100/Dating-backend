package config

import (
	"fmt"
	"log"
	"os"

	"model/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbUser := getEnv("DB_USER", "bee36693584e")
	dbPassword := getEnv("DB_PASSWORD", "Scout@1111")
	dbName := getEnv("DB_NAME", "state_model")
	dbPort := getEnv("DB_PORT", "3306")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&models.State{}, &models.Model{}, &models.FAQ{}, &models.GlobalPhone{})
	if err != nil {
		log.Printf("Failed to migrate database: %v", err)
		return nil, err
	}

	return db, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
