package usecase

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/internal/record/repository"
	"github.com/google/uuid"
)

type RecordService struct {
	recordRepo repository.RecordRepository
}

func NewRecordService(recordRepo repository.RecordRepository) RecordUseCase {
	return &RecordService{recordRepo: recordRepo}
}

func (s *RecordService) Save(record *entities.Record) error {
	return s.recordRepo.Save(record)
}

func (s *RecordService) FindByID(id uuid.UUID) (*entities.Record, error) {
	return s.recordRepo.FindByID(id)
}

func (s *RecordService) FindByUserID(user_id uuid.UUID, offset, limit int) ([]*entities.Record, int64, error) {
	return s.recordRepo.FindByUserID(user_id, offset, limit)
}

func (s *RecordService) Delete(id uuid.UUID) error {
	return s.recordRepo.Delete(id)
}
