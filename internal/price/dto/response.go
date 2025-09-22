package dto

type PriceResponse struct {
	UserID  string  `json:"user_id"`
	ID      string  `json:"id"`
	PrevID  string  `json:"prev_id"`
	NextID  string  `json:"next_id"`
	Amount  float32 `json:"amount"`
	BgColor string  `json:"bg_color"`
}
