package repository

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type CategoryRepository interface {
	Save(ctx context.Context, category *entities.Category) error
	FindByID(id uuid.UUID) (*entities.Category, error)
	FindByUserID(userID uuid.UUID) ([]*entities.Category, error)
	PatchValue(ctx context.Context, id uuid.UUID, category *entities.Category) error
	PatchPrev(ctx context.Context, id uuid.UUID, prevID *uuid.UUID) error
	PatchNext(ctx context.Context, id uuid.UUID, nextID *uuid.UUID) error
	Delete(id uuid.UUID) error
}
