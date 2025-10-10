package repository

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
)

type PriceRepository interface {
	Save(ctx context.Context, price *entities.Price) error
	FindByID(id string) (*entities.Price, error)
	FindByUserID(userID string) ([]*entities.Price, error)
	PatchValue(ctx context.Context, id string, price *entities.Price) error
	PatchPrev(ctx context.Context, id string, prevID string) error
	PatchNext(ctx context.Context, id string, nextID string) error
	Delete(id string) error
}
