package models

type AccountRequest struct {
	ID     int     `json:"id"`
	Income float64 `json:"income"`
}
