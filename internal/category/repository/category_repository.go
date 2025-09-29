package repository

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type CategoryRepository interface {
	Save(ctx context.Context, category *entities.Category) error
	FindByID(id uuid.UUID) (*entities.Category, error)
	FindByUserID(user_id uuid.UUID) ([]*entities.Category, error)
	Patch(ctx context.Context, id uuid.UUID, category *entities.Category) error
	Delete(id uuid.UUID) error
}
