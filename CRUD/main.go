package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/souravsk/go-zero-to-hero/config"
	"github.com/souravsk/go-zero-to-hero/handlers"
	"github.com/souravsk/go-zero-to-hero/models"
)

func main() {
	config.InitDB()

	// Run AutoMigrate to ensure the database schema is up to date
	if err := config.DB.AutoMigrate(&models.Todo{}); err != nil {
		fmt.Printf("Error during migration: %v\n", err)
		return
	}

	// Initialize the Gin router
	r := gin.Default()

	// Define your routes and handlers
	r.GET("/todos", handlers.GetTodos)
	r.POST("/todos", handlers.CreateTodo)
	r.GET("/todos/:id", handlers.GetTodoById)
	r.PUT("/todos/:id", handlers.UpdateTodo)
	r.DELETE("/todos/:id", handlers.DeleteTodo)

	// Start the server on port 8080
	r.Run(":8080")
}
