package dto

type PriceSaveRequest struct {
	UserID  string  `json:"user_id"`
	PrevID  string  `json:"prev_id"`
	NextID  string  `json:"next_id"`
	Amount  float32 `json:"amount"`
	BgColor string  `json:"bg_color"`
}

type PricePatchRequest struct {
	PrevID  string  `json:"prev_id"`
	NextID  string  `json:"next_id"`
	Amount  float32 `json:"amount"`
	BgColor string  `json:"bg_color"`
}
