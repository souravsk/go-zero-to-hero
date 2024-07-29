package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/souravsk/go-zero-to-hero/internal/database"
	"github.com/souravsk/go-zero-to-hero/internal/handler"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := database.Connect(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/todos", handler.GetTodos(db))
	r.POST("/todos", handler.CreateTodo(db))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
