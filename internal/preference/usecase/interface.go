package usecase

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
)

type PreferenceUseCase interface {
	FindByUserID(userID string) (*entities.Preference, error)
	Patch(userID string, preference *entities.Preference) (*entities.Preference, error)
}
