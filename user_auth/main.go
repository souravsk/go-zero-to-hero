package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/souravsk/go-zero-to-hero/user_auth/models"
	"github.com/souravsk/go-zero-to-hero/user_auth/routes"
)

func main() {
	r := gin.Default()

	// Fetch environment variables from Docker
	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize DB
	models.InitDB(config)

	// Load the router
	routes.AuthRoutes(r)

	// RUN the Server
	r.Run(":8080")
}
