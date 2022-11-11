package models

type TransferRequest struct {
	SenderID   int     `json:"sender_id"`
	ReceiverID int     `json:"receiver_id"`
	Amount     float64 `json:"amount"`
}
