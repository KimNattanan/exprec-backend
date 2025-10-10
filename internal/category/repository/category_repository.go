package repository

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
)

type CategoryRepository interface {
	Save(ctx context.Context, category *entities.Category) error
	FindByID(id string) (*entities.Category, error)
	FindByUserID(userID string) ([]*entities.Category, error)
	PatchValue(ctx context.Context, id string, category *entities.Category) error
	PatchPrev(ctx context.Context, id string, prevID string) error
	PatchNext(ctx context.Context, id string, nextID string) error
	Delete(id string) error
}
