package repository

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Save(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *GormUserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
