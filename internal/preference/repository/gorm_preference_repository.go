package repository

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"gorm.io/gorm"
)

type GormPreferenceRepository struct {
	db *gorm.DB
}

func NewGormPreferenceRepository(db *gorm.DB) PreferenceRepository {
	return &GormPreferenceRepository{db: db}
}

func (r *GormPreferenceRepository) FindByUserID(userID string) (*entities.Preference, error) {
	var preference entities.Preference
	if err := r.db.First(&preference, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &preference, nil
}

func (r *GormPreferenceRepository) Patch(userID string, preference *entities.Preference) error {
	result := r.db.Model(&entities.Preference{}).Where("user_id = ?", userID).Updates(preference)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
