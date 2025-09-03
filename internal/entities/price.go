package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Price struct {
	ID     uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	PrevID *uuid.UUID `gorm:"type:uuid" json:"prev_id"`
	NextID *uuid.UUID `gorm:"type:uuid" json:"next_id"`
	UserID uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	Amount int        `json:"amount"`

	Prev *Price `gorm:"foreignKey:PrevID;references:ID"`
	Next *Price `gorm:"foreignKey:NextID;references:ID"`
}

func (u *Price) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

func PreloadPrev(tx *gorm.DB) *gorm.DB {
	return tx.Preload("Next", PreloadPrev)
}
func PreloadNext(tx *gorm.DB) *gorm.DB {
	return tx.Preload("Next", PreloadNext)
}
