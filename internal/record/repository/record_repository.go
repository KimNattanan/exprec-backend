package repository

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type RecordRepository interface {
	Save(record *entities.Record) error
	CountByUserID(user_id uuid.UUID) (int64, error)
	FindByID(id uuid.UUID) (*entities.Record, error)
	FindByUserID(user_id uuid.UUID, offset, limit int) ([]*entities.Record, int64, error)
	Delete(id uuid.UUID) error
}
