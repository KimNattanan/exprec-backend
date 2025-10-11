package usecase

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type CategoryUseCase interface {
	Save(ctx context.Context, category *entities.Category) error
	FindByID(id uuid.UUID) (*entities.Category, error)
	FindByUserID(userID uuid.UUID) ([]*entities.Category, error)
	Patch(ctx context.Context, id uuid.UUID, category *entities.Category) (*entities.Category, error)
	Delete(id uuid.UUID) error
}
