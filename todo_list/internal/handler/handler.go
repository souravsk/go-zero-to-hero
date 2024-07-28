package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souravsk/go-zero-to-hero/internal/model"
)

func GetTodos(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title, completed FROM todos")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var todos []model.Todo
		for rows.Next() {
			var todo model.Todo
			if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			todos = append(todos, todo)
		}

		c.JSON(http.StatusOK, todos)
	}
}

func CreateTodo(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo model.Todo
		if err := c.BindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := `INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id`
		err := db.QueryRow(query, todo.Title, todo.Completed).Scan(&todo.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}
