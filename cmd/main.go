package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"hotai1806/crud-api/internal/todo" // Adjust import path to your module name
)

func main() {
	// Connect to Postgres (env vars from Docker)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	todo.DB = db

	// Auto-migrate schema
	db.AutoMigrate(&todo.Todo{})

	r := gin.Default()

	// Routes
	r.GET("/todos", todo.GetTodos)
	r.GET("/todos/:id", todo.GetTodo)
	r.POST("/todos", todo.CreateTodo)
	r.PUT("/todos/:id", todo.UpdateTodo)
	r.DELETE("/todos/:id", todo.DeleteTodo)

	r.Run(":8080")
}
