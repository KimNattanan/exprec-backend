package repository

import (
	"time"

	"github.com/KimNattanan/exprec-backend/internal/entities"
)

type RecordRepository interface {
	Save(record *entities.Record) error
	CountByUserID(userID string) (int64, error)
	FindByID(id string) (*entities.Record, error)
	FindByUserID(userID string, offset, limit int) ([]*entities.Record, error)
	FindByUserIDWithTimeRange(userID string, timeStart, timeEnd time.Time) ([]*entities.Record, error)
	Delete(id string) error
}
