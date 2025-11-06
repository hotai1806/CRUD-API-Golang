package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"todo-api/internal/todo" // Adjust to your module
	"todo-api/internal/user"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	todo.DB = db
	user.DB = db // Share DB

	// Auto-migrate
	db.AutoMigrate(&todo.Todo{}, &user.User{})

	r := gin.Default()

	// Public route
	r.POST("/login", user.Login)

	// Protected group
	protected := r.Group("/todos")
	protected.Use(todo.AuthMiddleware())
	{
		protected.GET("", todo.GetTodos)
		protected.GET("/:id", todo.GetTodo)
		protected.POST("", todo.CreateTodo)
		protected.PUT("/:id", todo.UpdateTodo)
		protected.DELETE("/:id", todo.DeleteTodo)
	}

	r.Run(":8080")
}
