package models

import "time"

type Account struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Balance   float64 `json:"balance"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
