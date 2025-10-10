package usecase

import (
	"time"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/internal/record/dto"
)

type RecordUseCase interface {
	Save(record *entities.Record) error
	FindByID(id string) (*entities.Record, error)
	FindByUserID(userID string, offset, limit int) ([]*entities.Record, int64, error)
	Delete(id string) error
	GetDashboardDataByUserID(userID string, dateFirst, dateLast time.Time) (*dto.DashboardData, error)
}
