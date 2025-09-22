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
	Patch(ctx context.Context, id uuid.UUID, price *entities.Price) error
	Delete(id uuid.UUID) error
}
