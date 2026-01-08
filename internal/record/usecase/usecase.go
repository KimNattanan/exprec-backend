package usecase

import (
	"time"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/internal/record/dto"
	"github.com/KimNattanan/exprec-backend/internal/record/repository"
	"github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/google/uuid"
)

type RecordService struct {
	recordRepo repository.RecordRepository
}

func NewRecordService(recordRepo repository.RecordRepository) RecordUseCase {
	return &RecordService{recordRepo: recordRepo}
}

func (s *RecordService) Save(record *entities.Record) error {
	cnt, err := s.recordRepo.CountByUserID(record.UserID)
	if err != nil {
		return err
	}
	if cnt >= 100 {
		return apperror.ErrLimitExceeded
	}
	return s.recordRepo.Save(record)
}

func (s *RecordService) FindByID(id uuid.UUID) (*entities.Record, error) {
	return s.recordRepo.FindByID(id)
}

func (s *RecordService) FindByUserID(userID uuid.UUID, offset, limit int) ([]*entities.Record, int64, error) {
	records, err := s.recordRepo.FindByUserID(userID, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	totalRecords, err := s.recordRepo.CountByUserID(userID)
	if err != nil {
		return nil, 0, err
	}
	return records, totalRecords, nil
}

func (s *RecordService) GetDashboardDataByUserID(userID uuid.UUID, timeStart time.Time, timeEnd time.Time) (*dto.DashboardData, error) {
	records, err := s.recordRepo.FindByUserIDWithTimeRange(userID, timeStart, timeEnd)
	if err != nil {
		return nil, err
	}

	var totalAmount float32
	amountByCategory := make(map[string]float32)
	categoryColor := make(map[string]string)
	for _, e := range records {
		totalAmount += e.Amount
		amountByCategory[e.Category] += e.Amount
		categoryColor[e.Category] = e.CategoryBgColor
	}

	return &dto.DashboardData{
		TotalAmount:      totalAmount,
		AmountByCategory: amountByCategory,
		CategoryColor:    categoryColor,
		Records:          records,
	}, nil
}

func (s *RecordService) Delete(id uuid.UUID) error {
	return s.recordRepo.Delete(id)
}
