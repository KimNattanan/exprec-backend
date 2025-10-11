package usecase

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type PreferenceUseCase interface {
	FindByUserID(userID uuid.UUID) (*entities.Preference, error)
	Patch(userID uuid.UUID, preference *entities.Preference) (*entities.Preference, error)
}
