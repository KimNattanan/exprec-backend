package usecase

import (
	"errors"
	"os"
	"time"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/internal/user/repository"
	"github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserUseCase {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *entities.User) error {
	existingUser, err := s.repo.FindByEmail(user.Email)
	if existingUser != nil {
		return apperror.ErrAlreadyExists
	}
	if !errors.Is(err, apperror.ErrRecordNotFound) {
		return err
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(pw)

	return s.repo.Save(user)
}

func (s *UserService) Login(email, password string) (string, *entities.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || user == nil {
		return "", nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, err
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", nil, err
	}
	return t, user, nil
}
