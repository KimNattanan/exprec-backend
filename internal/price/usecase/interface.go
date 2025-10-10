package usecase

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
)

type PriceUseCase interface {
	Save(ctx context.Context, price *entities.Price) error
	FindByID(id string) (*entities.Price, error)
	FindByUserID(userID string) ([]*entities.Price, error)
	Patch(ctx context.Context, id string, price *entities.Price) (*entities.Price, error)
	Delete(id string) error
}
