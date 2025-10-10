package repository

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
)

type UserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	FindByID(id string) (*entities.User, error)
	FindAll() ([]*entities.User, error)
	Save(user *entities.User) error
	Patch(id string, user *entities.User) error
	Delete(id string) error
}
