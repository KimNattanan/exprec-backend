package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Record struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID     uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	CreatedAt  time.Time  `json:"created_at"`
	Amount     float32    `json:"amount"`
	CategoryID *uuid.UUID `gorm:"type:uuid" json:"category_id"`
	Note       string     `gorm:"type:varchar(255)" json:"note"`

	Category Category `gorm:"foreignKey:CategoryID;reference:ID;constraint:OnDelete:SET NULL"`
}

func (u *Record) BeforeCreate(d *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	return
}
