package repository

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"gorm.io/gorm"
)

type GormSessionRepository struct {
	db *gorm.DB
}

func NewGormSessionRepository(db *gorm.DB) SessionRepository {
	return &GormSessionRepository{db: db}
}

func (r *GormSessionRepository) Save(session *entities.Session) error {
	return r.db.Create(session).Error
}

func (r *GormSessionRepository) FindByID(id string) (*entities.Session, error) {
	var session entities.Session
	if err := r.db.First(&session, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *GormSessionRepository) Revoke(id string) error {
	result := r.db.Model(&entities.Session{}).Where("id = ?", id).Update("is_revoked", true)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormSessionRepository) Delete(id string) error {
	var session entities.Session
	if err := r.db.First(&session, "id = ?", id).Error; err != nil {
		return err
	}
	result := r.db.Delete(&session)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
