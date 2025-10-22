package usecase

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/internal/session/repository"
)

type SessionService struct {
	sessionRepo repository.SessionRepository
}

func NewSessionService(sessionRepo repository.SessionRepository) SessionUseCase {
	return &SessionService{
		sessionRepo: sessionRepo,
	}
}

func (s *SessionService) Save(session *entities.Session) error {
	return s.sessionRepo.Save(session)
}
func (s *SessionService) FindByID(id string) (*entities.Session, error) {
	return s.sessionRepo.FindByID(id)
}
func (s *SessionService) Revoke(id string) error {
	return s.sessionRepo.Revoke(id)
}
func (s *SessionService) Delete(id string) error {
	return s.sessionRepo.Delete(id)
}
