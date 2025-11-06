package todo

import (
	"gorm.io/gorm"
)

// Todo model
type Todo struct {
	gorm.Model        // Embeds ID, CreatedAt, UpdatedAt, DeletedAt
	Task       string `json:"task"`
	Done       bool   `json:"done"`
}

// DB is the global DB connection (set in main)
var DB *gorm.DB

// GetAll fetches all todos
func GetAll() ([]Todo, error) {
	var todos []Todo
	err := DB.Find(&todos).Error
	return todos, err
}

// GetByID fetches todo by ID
func GetByID(id uint) (Todo, error) {
	var todo Todo
	err := DB.First(&todo, id).Error
	return todo, err
}

// Create adds a new todo
func Create(todo *Todo) error {
	return DB.Create(todo).Error
}

// Update modifies a todo
func Update(todo *Todo) error {
	return DB.Save(todo).Error
}

// Delete removes a todo by ID
func Delete(id uint) error {
	return DB.Delete(&Todo{}, id).Error
}
