package models

import "time"

type RequestStatus string

const (
	accepted RequestStatus = "Accepted"
	pending  RequestStatus = "Pending"
	declined RequestStatus = "Declined"
	canceled RequestStatus = "Canceled"
)

type Request struct {
	ID          uint          `json:"id" gorm:"primaryKey"`
	UserId      uint          `json:"userId"`
	Status      RequestStatus `json:"status"`
	Description string        `json:"description"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	Services    []Service     `gorm:"many2many:request-services;"`
}
