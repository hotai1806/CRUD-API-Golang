package booking

import (
	"time"

	"gorm.io/gorm"
)

// Booking model (simple example)
type Booking struct {
	gorm.Model
	UserID  uint      `json:"user_id"` // Tie to user
	Date    time.Time `json:"date"`
	Status  string    `json:"status"` // e.g., "pending", "confirmed"
	Details string    `json:"details"`
}

// DB shared
var DB *gorm.DB

// GetAll fetches all bookings (filter by user in handler for security)
func GetAll() ([]Booking, error) {
	var bookings []Booking
	err := DB.Find(&bookings).Error
	return bookings, err
}

// GetByID fetches booking by ID
func GetByID(id uint) (Booking, error) {
	var booking Booking
	err := DB.First(&booking, id).Error
	return booking, err
}

// Create adds a new booking
func Create(booking *Booking) error {
	return DB.Create(booking).Error
}

// Update modifies a booking
func Update(booking *Booking) error {
	return DB.Save(booking).Error
}

// Delete removes a booking by ID
func Delete(id uint) error {
	return DB.Delete(&Booking{}, id).Error
}
