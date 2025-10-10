package repository

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"gorm.io/gorm"
)

type GormCategoryRepository struct {
	db *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) CategoryRepository {
	return &GormCategoryRepository{db: db}
}

func (r *GormCategoryRepository) Save(ctx context.Context, category *entities.Category) error {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		tx = r.db
	}
	return tx.Create(category).Error
}

func (r *GormCategoryRepository) FindByID(id string) (*entities.Category, error) {
	var category entities.Category
	if err := r.db.Preload("Prev").Preload("Next").First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
func (r *GormCategoryRepository) FindByUserID(userID string) ([]*entities.Category, error) {
	var categoryValues []entities.Category
	if err := r.db.Preload("Prev").Preload("Next").Find(&categoryValues, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	categories := make([]*entities.Category, len(categoryValues))
	for i := range categories {
		categories[i] = &categoryValues[i]
	}
	return categories, nil
}

func (r *GormCategoryRepository) PatchValue(ctx context.Context, id string, category *entities.Category) error {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		tx = r.db
	}
	result := tx.Model(&entities.Category{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title":    category.Title,
		"bg_color": category.BgColor,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (r *GormCategoryRepository) PatchPrev(ctx context.Context, id string, prevID string) error {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		tx = r.db
	}
	result := tx.Model(&entities.Category{}).Where("id = ?", id).Update("prev_id", prevID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (r *GormCategoryRepository) PatchNext(ctx context.Context, id string, nextID string) error {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		tx = r.db
	}
	result := tx.Model(&entities.Category{}).Where("id = ?", id).Update("next_id", nextID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormCategoryRepository) Delete(id string) error {
	var category entities.Category
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return err
	}
	result := r.db.Delete(&category)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
