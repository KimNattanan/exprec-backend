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

func PreloadPrices(depth int) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if depth <= 0 {
			return d
		}
		return d.Preload("Prices", PreloadPrices(depth-1))
	}
}
