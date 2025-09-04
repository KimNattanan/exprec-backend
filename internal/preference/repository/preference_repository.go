package repository

import "github.com/KimNattanan/exprec-backend/internal/entities"

type PreferenceRepository interface {
	Save(pref *entities.Preference) error
	// FindByID(id string) (*entities.User, error)
	// FindAll() ([]*entities.User, error)
	// Patch(id string, user *entities.User) error
	// Delete(id string) error
}
