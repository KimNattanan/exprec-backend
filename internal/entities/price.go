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

func PreloadPrev(depth int) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if depth <= 0 {
			return d
		}
		return d.Preload("Prev", PreloadPrev(depth-1))
	}
}

func PreloadNext(depth int) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if depth <= 0 {
			return d
		}
		return d.Preload("Prev", PreloadNext(depth-1))
	}
}
