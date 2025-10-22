package usecase

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
)

type SessionUseCase interface {
	Save(session *entities.Session) error
	FindByID(id string) (*entities.Session, error)
	Revoke(id string) error
	Delete(id string) error
}
