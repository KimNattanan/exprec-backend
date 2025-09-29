package repository

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type RecordRepository interface {
	Save(record *entities.Record) error
	FindByID(id uuid.UUID) (*entities.Record, error)
	FindByUserID(user_id uuid.UUID) ([]*entities.Record, error)
	Delete(id uuid.UUID) error
}
