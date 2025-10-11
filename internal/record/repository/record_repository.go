package repository

import (
	"time"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type RecordRepository interface {
	Save(record *entities.Record) error
	CountByUserID(userID uuid.UUID) (int64, error)
	FindByID(id uuid.UUID) (*entities.Record, error)
	FindByUserID(userID uuid.UUID, offset, limit int) ([]*entities.Record, error)
	FindByUserIDWithTimeRange(userID uuid.UUID, timeStart, timeEnd time.Time) ([]*entities.Record, error)
	Delete(id uuid.UUID) error
}
