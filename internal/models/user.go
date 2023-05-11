package models

import "time"

type User struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	Surname       string    `json:"surname"`
	LastName      string    `json:"lastname"`
	Email         string    `json:"email"`
	Password      string    `json:"-"`
	Phone         string    `json:"phone"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	IsAdmin       bool      `json:"isAdmin"`
	Requests      []Request
	Notifications []Notification
}
