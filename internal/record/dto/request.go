package dto

type RecordSaveRequest struct {
	UserID          string  `json:"user_id"`
	Amount          float32 `json:"amount"`
	Category        string  `json:"category"`
	CategoryBgColor string  `json:"category_bg_color"`
	Note            string  `json:"note"`
}
