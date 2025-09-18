package dto

import "github.com/google/uuid"

type PriceResponse struct {
	ID      uuid.UUID `json:"id"`
	PrevID  uuid.UUID `json:"prev_id"`
	NextID  uuid.UUID `json:"next_id"`
	UserID  uuid.UUID `json:"user_id"`
	Amount  float32   `json:"amount"`
	BgColor string    `json:"bg_color"`
}
