package repository

import "github.com/KimNattanan/exprec-backend/internal/entities"

type UserRepository interface {
	Save(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	Patch(id string, user *entities.User) error
	Delete(id string) error
	// FindByID(id string) (*entities.User, error)
	// FindAll() ([]*entities.User, error)
}
