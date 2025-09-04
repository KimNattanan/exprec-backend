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

	Prices     []Price    `gorm:"foreignKey:UserID"`
	Categories []Category `gorm:"foreignKey:UserID"`
	Records    []Record   `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(d *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
