package dto

import (
	"time"

	"github.com/KimNattanan/exprec-backend/internal/entities"
)

type RecordResponse struct {
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	Amount          float32   `json:"amount"`
	Category        string    `json:"category"`
	CategoryBgColor string    `json:"category_bg_color"`
	Note            string    `json:"note"`
}

type DashboardData struct {
	TotalAmount      float32
	AmountByCategory map[string]float32
	Records          []*entities.Record
}
