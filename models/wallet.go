package models

import "time"

type Wallet struct {
	ID       int       `json:"id"`
	Amount   float64   `json:"amount"`
	Category string    `json:"category"`
	AddedAt  time.Time `json:"added_at"`
}
