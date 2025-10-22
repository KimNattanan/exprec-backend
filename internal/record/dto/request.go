package dto

type RecordSaveRequest struct {
	CreatedAt       string  `json:"created_at"`
	Amount          float32 `json:"amount"`
	Category        string  `json:"category"`
	CategoryBgColor string  `json:"category_bg_color"`
	Note            string  `json:"note"`
}
