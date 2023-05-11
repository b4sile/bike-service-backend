package models

type Service struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Price       float64 `json:"price"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}
