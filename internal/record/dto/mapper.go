package dto

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

func ToRecordResponse(record *entities.Record) *RecordResponse {
	return &RecordResponse{
		UserID:   record.UserID.String(),
		ID:       record.ID.String(),
		Category: record.Category,
		Amount:   record.Amount,
		Note:     record.Note,
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
	userID, err := uuid.Parse(record.UserID)
	if err != nil {
		return nil, err
	}
	return &entities.Record{
		UserID:   userID,
		Category: record.Category,
		Amount:   record.Amount,
		Note:     record.Note,
	}, nil
}
