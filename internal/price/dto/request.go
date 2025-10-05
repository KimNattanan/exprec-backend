package dto

type PriceSaveRequest struct {
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
