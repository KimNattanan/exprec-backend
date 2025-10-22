package dto

import (
	"time"

	"github.com/KimNattanan/exprec-backend/internal/entities"
)

func ToRecordResponse(record *entities.Record) *RecordResponse {
	return &RecordResponse{
		ID:              record.ID.String(),
		CreatedAt:       record.CreatedAt,
		Amount:          record.Amount,
		Category:        record.Category,
		CategoryBgColor: record.CategoryBgColor,
		Note:            record.Note,
	}
}

func ToRecordResponseList(records []*entities.Record) []*RecordResponse {
	result := make([]*RecordResponse, len(records))
	for i, u := range records {
		result[i] = ToRecordResponse(u)
	}
	return result
}

func FromRecordSaveRequest(record *RecordSaveRequest) (*entities.Record, error) {
	createdAt, err := time.Parse(time.RFC3339, record.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &entities.Record{
		CreatedAt:       createdAt,
		Amount:          record.Amount,
		Category:        record.Category,
		CategoryBgColor: record.CategoryBgColor,
		Note:            record.Note,
	}, nil
}
