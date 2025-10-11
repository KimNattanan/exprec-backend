package repository

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	FindByID(id uuid.UUID) (*entities.User, error)
	FindAll() ([]*entities.User, error)
	Save(user *entities.User) error
	Patch(id uuid.UUID, user *entities.User) error
	Delete(id uuid.UUID) error
}
