package usecase

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type PriceUseCase interface {
	Save(ctx context.Context, price *entities.Price) error
	FindByID(id uuid.UUID) (*entities.Price, error)
	FindByUserID(userID uuid.UUID) ([]*entities.Price, error)
	Patch(ctx context.Context, id uuid.UUID, price *entities.Price) (*entities.Price, error)
	Delete(id uuid.UUID) error
}
