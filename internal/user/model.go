package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
}

// DB is shared from todo package (or make it global)
var DB *gorm.DB

// CreateUser adds a new user with hashed password
func CreateUser(u *User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return DB.Create(u).Error
}

// FindByUsername finds user by username
func FindByUsername(username string) (User, error) {
	var u User
	err := DB.Where("username = ?", username).First(&u).Error
	return u, err
}
