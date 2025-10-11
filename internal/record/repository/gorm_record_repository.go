package repository

import (
	"time"

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

func (r *GormRecordRepository) CountByUserID(userID uuid.UUID) (int64, error) {
	var cnt int64
	if err := r.db.Model(&entities.Record{}).Where("user_id = ?", userID).Count(&cnt).Error; err != nil {
		return 0, err
	}
	return cnt, nil
}

func (r *GormRecordRepository) FindByID(id uuid.UUID) (*entities.Record, error) {
	var record entities.Record
	if err := r.db.First(&record, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}
func (r *GormRecordRepository) FindByUserID(userID uuid.UUID, offset, limit int) ([]*entities.Record, error) {
	var recordValues []entities.Record
	if err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Offset(offset).Limit(limit).Find(&recordValues).Error; err != nil {
		return nil, err
	}
	records := make([]*entities.Record, len(recordValues))
	for i := range records {
		records[i] = &recordValues[i]
	}
	return records, nil
}

func (r *GormRecordRepository) FindByUserIDWithTimeRange(userID uuid.UUID, timeStart, timeEnd time.Time) ([]*entities.Record, error) {
	var recordValues []entities.Record
	if err := r.db.Where("user_id = ? AND created_at >= ? AND created_at <= ?", userID, timeStart, timeEnd).Order("created_at DESC").Find(&recordValues).Error; err != nil {
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
