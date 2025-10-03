package repository

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type PriceRepository interface {
	Save(ctx context.Context, price *entities.Price) error
	FindByID(id uuid.UUID) (*entities.Price, error)
	FindByUserID(user_id uuid.UUID) ([]*entities.Price, error)
	PatchValue(ctx context.Context, id uuid.UUID, price *entities.Price) error
	PatchPrev(ctx context.Context, id uuid.UUID, prevID *uuid.UUID) error
	PatchNext(ctx context.Context, id uuid.UUID, nextID *uuid.UUID) error
	Delete(id uuid.UUID) error
}
