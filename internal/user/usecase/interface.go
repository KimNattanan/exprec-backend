package usecase

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type UserUseCase interface {
	Register(user *entities.User) error
	Login(email, password string) (string, *entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindByID(id uuid.UUID) (*entities.User, error)
	FindAll() ([]*entities.User, error)
	Patch(id uuid.UUID, user *entities.User) (*entities.User, error)
	Delete(id uuid.UUID) error
}
