package models

import (
	"os"
	"time"

	"github.com/b4sile/bike-service-backend/internal/helpers"
	"gorm.io/gorm"
)

type User struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Name          string         `json:"name"`
	Surname       string         `json:"surname"`
	Lastname      string         `json:"lastname"`
	Email         string         `json:"email" gorm:"uniqueIndex"`
	Password      string         `json:"-"`
	Phone         string         `json:"phone"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	IsAdmin       bool           `json:"isAdmin"`
	Requests      []Request      `json:"-"`
	Notifications []Notification `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Password = helpers.HashPassword(u.Password)
	return nil
}

func CreateAdminUser() error {
	var user User
	if err := DB.First(&user, "is_admin = ?", true).Error; err != nil {
		user.Email = os.Getenv("ADMIN_EMAIL")
		user.Password = os.Getenv("ADMIN_PASSWORD")
		user.IsAdmin = true
		DB.Create(&user)
		return nil
	}
	return nil
}
