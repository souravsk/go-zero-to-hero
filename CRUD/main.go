package main

import (
	"github.com/gin-gonic/gin"
	"github.com/souravsk/go-zero-to-hero/config"
	"github.com/souravsk/go-zero-to-hero/handlers"
	"github.com/souravsk/go-zero-to-hero/models"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	config.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()

	r.GET("/todos", handlers.GetTodos)
	r.POST("/todos", handlers.CreateTodo)
	r.GET("/todos/:id", handlers.GetTodoById)
	r.PUT("/todos/:id", handlers.UpdateTodo)
	r.DELETE("/todos/:id", handlers.DeleteTodo)

	r.Run(":8080")
}
