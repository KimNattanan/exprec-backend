package repository

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormRecordRepository struct {
	db *gorm.DB
}

func NewGormRecordRepository(db *gorm.DB) RecordRepository {
	return &GormRecordRepository{db: db}
}

func (r *GormRecordRepository) Save(record *entities.Record) error {
	return r.db.Create(record).Error
}

func (r *GormRecordRepository) FindByID(id uuid.UUID) (*entities.Record, error) {
	var record entities.Record
	if err := r.db.First(&record, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}
func (r *GormRecordRepository) FindByUserID(user_id uuid.UUID) ([]*entities.Record, error) {
	var recordValues []entities.Record
	if err := r.db.Find(&recordValues, "user_id = ?", user_id).Error; err != nil {
		return nil, err
	}
	records := make([]*entities.Record, len(recordValues))
	for i := range records {
		records[i] = &recordValues[i]
	}
	return records, nil
}

func (r *GormRecordRepository) Delete(id uuid.UUID) error {
	var record entities.Record
	if err := r.db.First(&record, "id = ?", id).Error; err != nil {
		return err
	}
	result := r.db.Delete(&record)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
