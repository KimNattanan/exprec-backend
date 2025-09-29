package dto

type RecordSaveRequest struct {
	UserID   string  `json:"user_id"`
	Amount   float32 `json:"amount"`
	Category string  `json:"category"`
	Note     string  `json:"note"`
}
