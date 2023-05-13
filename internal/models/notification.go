package models

import "time"

type Notification struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserId      uint      `json:"userId"`
	IsReaded    bool      `json:"isReaded"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}
