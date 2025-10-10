package usecase

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
)

type CategoryUseCase interface {
	Save(ctx context.Context, category *entities.Category) error
	FindByID(id string) (*entities.Category, error)
	FindByUserID(userID string) ([]*entities.Category, error)
	Patch(ctx context.Context, id string, category *entities.Category) (*entities.Category, error)
	Delete(id string) error
}
