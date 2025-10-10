package repository

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
)

type PreferenceRepository interface {
	FindByUserID(userID string) (*entities.Preference, error)
	Patch(userID string, preference *entities.Preference) error
}
