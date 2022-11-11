package models

import "time"

type Transaction struct {
	SenderID   int
	ReceiverID int
	Amount     float64
	CreatedAt  time.Time
}
