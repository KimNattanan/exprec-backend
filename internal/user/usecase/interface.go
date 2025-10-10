package usecase

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"golang.org/x/oauth2"
)

type UserUseCase interface {
	Register(user *entities.User) error
	Login(email, password string) (string, *entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindByID(id string) (*entities.User, error)
	FindAll() ([]*entities.User, error)
	Patch(id string, user *entities.User) (*entities.User, error)
	Delete(id string) error
	LoginOrRegisterWithGoogle(userInfo map[string]interface{}, oauthToken *oauth2.Token) (string, *entities.User, error)
}
