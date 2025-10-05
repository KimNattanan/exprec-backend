package dto

type CategorySaveRequest struct {
	PrevID  string `json:"prev_id"`
	NextID  string `json:"next_id"`
	Title   string `json:"title"`
	BgColor string `json:"bg_color"`
}

type CategoryPatchRequest struct {
	PrevID  string `json:"prev_id"`
	NextID  string `json:"next_id"`
	Title   string `json:"title"`
	BgColor string `json:"bg_color"`
}
