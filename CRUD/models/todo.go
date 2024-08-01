package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"` // Change to bool if you want true/false instead of a string
}

func GetAllTodos(db *gorm.DB) ([]Todo, error) {
	var todos []Todo
	if err := db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func CreateTodo(db *gorm.DB, todo *Todo) error {
	if err := db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodoById(db *gorm.DB, id uint) (*Todo, error) {
	var todo Todo
	if err := db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func UpdateTodo(db *gorm.DB, todo *Todo) error {
	if err := db.Save(todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTodo(db *gorm.DB, id uint) error {
	if err := db.Delete(&Todo{}, id).Error; err != nil {
		return err
	}
	return nil
}
