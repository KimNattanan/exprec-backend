package repository

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormPriceRepository struct {
	db *gorm.DB
}

func NewGormPriceRepository(db *gorm.DB) PriceRepository {
	return &GormPriceRepository{db: db}
}

func (r *GormPriceRepository) Save(ctx context.Context, price *entities.Price) error {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		tx = r.db
	}
	return tx.Create(price).Error
}

func (r *GormPriceRepository) FindByID(id uuid.UUID) (*entities.Price, error) {
	var price entities.Price
	if err := r.db.Preload("Prev").Preload("Next").First(&price, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &price, nil
}

func (r *GormPriceRepository) Patch(ctx context.Context, id uuid.UUID, price *entities.Price) error {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		tx = r.db
	}
	result := tx.Model(&entities.Price{}).Where("id = ?", id).Updates(price)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormPriceRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&entities.Price{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
