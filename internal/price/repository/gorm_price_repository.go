package repository

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
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

func (r *GormPriceRepository) FindByID(id string) (*entities.Price, error) {
	var price entities.Price
	if err := r.db.Preload("Prev").Preload("Next").First(&price, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &price, nil
}
func (r *GormPriceRepository) FindByUserID(userID string) ([]*entities.Price, error) {
	var priceValues []entities.Price
	if err := r.db.Preload("Prev").Preload("Next").Find(&priceValues, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	prices := make([]*entities.Price, len(priceValues))
	for i := range prices {
		prices[i] = &priceValues[i]
	}
	return prices, nil
}

func (r *GormPriceRepository) PatchValue(ctx context.Context, id string, price *entities.Price) error {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		tx = r.db
	}
	result := tx.Model(&entities.Price{}).Where("id = ?", id).Updates(map[string]interface{}{
		"amount":   price.Amount,
		"bg_color": price.BgColor,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (r *GormPriceRepository) PatchPrev(ctx context.Context, id string, prevID string) error {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		tx = r.db
	}
	result := tx.Model(&entities.Price{}).Where("id = ?", id).Update("prev_id", prevID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (r *GormPriceRepository) PatchNext(ctx context.Context, id string, nextID string) error {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		tx = r.db
	}
	result := tx.Model(&entities.Price{}).Where("id = ?", id).Update("next_id", nextID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormPriceRepository) Delete(id string) error {
	var price entities.Price
	if err := r.db.First(&price, "id = ?", id).Error; err != nil {
		return err
	}
	result := r.db.Delete(&price)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
