package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"hotai1806/crud-api/internal/booking"
	"hotai1806/crud-api/internal/middleware"
	"hotai1806/crud-api/internal/todo" // Adjust to your module
	"hotai1806/crud-api/internal/user"
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

	// Protected group
	// Group all API routes
	api := r.Group("/api/v1")

	// Apply auth middleware to all protected routes
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	// Public route
	api.POST("/login", user.Login)
	api.POST("/create-user", user.NewUser)

	// Todos routes
	todos := protected.Group("/todos")
	{
		todos.GET("", todo.GetTodos)
		todos.GET("/:id", todo.GetTodo)
		todos.POST("", todo.CreateTodo)
		todos.PUT("/:id", todo.UpdateTodo)
		todos.DELETE("/:id", todo.DeleteTodo)
	}

	// Bookings routes
	bookings := protected.Group("/bookings")
	{
		bookings.GET("", booking.GetBookings)
		bookings.GET("/:id", booking.GetBooking)
		bookings.POST("", booking.CreateBooking)
		bookings.DELETE("/:id", booking.DeleteBooking)
	}

	r.Run(":8080")
}
