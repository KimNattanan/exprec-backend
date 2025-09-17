package usecase

import (
	"errors"
	"os"
	"time"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/internal/user/repository"
	"github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserUseCase {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(user *entities.User) error {
	existingUser, err := s.userRepo.FindByEmail(user.Email)
	if existingUser != nil {
		return apperror.ErrAlreadyExists
	}
	if !errors.Is(err, apperror.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.Save(user)
}

func (s *UserService) Login(email, password string) (string, *entities.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil || user == nil {
		return "", nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, err
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // 3 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}

func (s *UserService) FindByEmail(email string) (*entities.User, error) {
	return s.userRepo.FindByEmail(email)
}

func (s *UserService) FindByID(id uuid.UUID) (*entities.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) FindAll() ([]*entities.User, error) {
	return s.userRepo.FindAll()
}

func (s *UserService) Patch(id uuid.UUID, user *entities.User) (*entities.User, error) {
	if err := s.userRepo.Patch(id, user); err != nil {
		return nil, err
	}
	return s.userRepo.FindByID(id)
}

func (s *UserService) Delete(id uuid.UUID) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) LoginOrRegisterWithGoogle(userInfo map[string]interface{}, token *oauth2.Token) (string, *entities.User, error) {
	email, ok := userInfo["email"].(string)
	name, _ := userInfo["name"].(string)
	if !ok || email == "" {
		return "", nil, apperror.ErrInvalidData
	}
	user, err := s.userRepo.FindByEmail(email)
	if err != nil && !errors.Is(err, apperror.ErrRecordNotFound) {
		return "", nil, err
	}
	if user == nil {
		user = &entities.User{
			Email:    email,
			Name:     name,
			Password: "",
		}
		if err := s.userRepo.Save(user); err != nil {
			return "", nil, err
		}
	}
	user.Password = ""

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // 3 days
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_SECRET")
	tokenString, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", nil, err
	}
	return tokenString, user, nil
}
