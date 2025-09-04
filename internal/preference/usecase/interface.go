package usecase

import "github.com/KimNattanan/exprec-backend/internal/entities"

type UserUseCase interface {
	Register(user *entities.User) error
	Login(email, password string) (string, *entities.User, error)
}
