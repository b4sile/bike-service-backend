package models

import "time"

type RequestStatusType string

const (
	accepted RequestStatusType = "Accepted"
	pending  RequestStatusType = "Pending"
	declined RequestStatusType = "Declined"
	canceled RequestStatusType = "Canceled"
)

var RequestStatus = struct {
	Accepted RequestStatusType
	Pending  RequestStatusType
	Declined RequestStatusType
	Canceled RequestStatusType
}{Accepted: accepted, Pending: pending, Declined: declined, Canceled: canceled}

type Request struct {
	ID             uint              `json:"id" gorm:"primaryKey"`
	UserId         uint              `json:"userId"`
	Status         RequestStatusType `json:"status"`
	Description    string            `json:"description"`
	Phone          string            `json:"phone"`
	Latitude       float64           `json:"latitude"`
	Longitude      float64           `json:"longitude"`
	IsOuterRequest bool              `json:"isOuterRequest"`
	CreatedAt      time.Time         `json:"createdAt"`
	UpdatedAt      time.Time         `json:"updatedAt"`
	Services       []Service         `json:"-" gorm:"many2many:request-services;"`
}
