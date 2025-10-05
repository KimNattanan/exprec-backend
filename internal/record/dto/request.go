package dto

type RecordSaveRequest struct {
	Amount          float32 `json:"amount"`
	Category        string  `json:"category"`
	CategoryBgColor string  `json:"category_bg_color"`
	Note            string  `json:"note"`
}
