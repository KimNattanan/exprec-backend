package dto

import (
	"time"
)

type RecordResponse struct {
	UserID          string    `json:"user_id"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	Amount          float32   `json:"amount"`
	Category        string    `json:"category"`
	CategoryBgColor string    `json:"category_bg_color"`
	Note            string    `json:"note"`
}
