package usecase

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/internal/preference/repository"
	"github.com/google/uuid"
)

type PreferenceService struct {
	repo repository.PreferenceRepository
}

func NewPreferenceService(repo repository.PreferenceRepository) PreferenceUseCase {
	return &PreferenceService{repo: repo}
}

func (s *PreferenceService) FindByUserID(userID uuid.UUID) (*entities.Preference, error) {
	return s.repo.FindByUserID(userID)
}

func (s *PreferenceService) Save(preference *entities.Preference) error {
	return s.repo.Save(preference)
}

func (s *PreferenceService) Patch(userID uuid.UUID, preference *entities.Preference) (*entities.Preference, error) {
	if err := s.repo.Patch(userID, preference); err != nil {
		return nil, err
	}
	return s.repo.FindByUserID(userID)
}
