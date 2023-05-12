package models

import "time"

type User struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Name          string         `json:"name"`
	Surname       string         `json:"surname"`
	Lastname      string         `json:"lastname"`
	Email         string         `json:"email"`
	Password      string         `json:"-"`
	Phone         string         `json:"phone"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	IsAdmin       bool           `json:"isAdmin"`
	Requests      []Request      `json:"-"`
	Notifications []Notification `json:"-"`
}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
