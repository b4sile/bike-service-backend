package models

import (
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
	Password      string         `json:"password"`
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
