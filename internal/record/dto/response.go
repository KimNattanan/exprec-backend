package dto

import (
	"time"
)

type RecordResponse struct {
	UserID    string    `json:"user_id"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Category  string    `json:"category"`
	Amount    float32   `json:"amount"`
	Note      string    `json:"note"`
}
