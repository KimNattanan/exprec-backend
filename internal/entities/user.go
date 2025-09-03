package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email    string    `gorm:"uniqueIndex" json:"email"`
	Password string    `json:"password"`
	Name     string    `json:"name"`

	Prices []Price `gorm:"foreignKey:UserID" json:"prices"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

func PreloadPrices(tx *gorm.DB) *gorm.DB {
	return tx.Preload("Prices", PreloadPrices)
}
