package dto

type CategoryResponse struct {
	ID      string `json:"id"`
	PrevID  string `json:"prev_id"`
	NextID  string `json:"next_id"`
	Title   string `json:"title"`
	BgColor string `json:"bg_color"`
}
