package usecase

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type RecordUseCase interface {
	Save(record *entities.Record) error
	FindByID(id uuid.UUID) (*entities.Record, error)
	FindByUserID(user_id uuid.UUID) ([]*entities.Record, error)
	Delete(id uuid.UUID) error
}
