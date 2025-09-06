package repository

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type PreferenceRepository interface {
	FindByUserID(userID uuid.UUID) (*entities.Preference, error)
	Save(preference *entities.Preference) error
	Patch(userID uuid.UUID, preference *entities.Preference) error
}
