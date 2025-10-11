package usecase

import (
	"time"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/internal/record/dto"
	"github.com/google/uuid"
)

type RecordUseCase interface {
	Save(record *entities.Record) error
	FindByID(id uuid.UUID) (*entities.Record, error)
	FindByUserID(userID uuid.UUID, offset, limit int) ([]*entities.Record, int64, error)
	GetDashboardDataByUserID(userID uuid.UUID, dateFirst, dateLast time.Time) (*dto.DashboardData, error)
	Delete(id uuid.UUID) error
}
