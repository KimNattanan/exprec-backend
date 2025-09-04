package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID      uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	PrevID  *uuid.UUID `gorm:"type:uuid" json:"prev_id"`
	NextID  *uuid.UUID `gorm:"type:uuid" json:"next_id"`
	UserID  uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	Title   string     `gorm:"type:varchar(50)" json:"title"`
	BgColor string     `gorm:"type:char(7)" json:"bg_color"`

	Prev *Category `gorm:"foreignKey:PrevID;references:ID"`
	Next *Category `gorm:"foreignKey:NextID;references:ID"`
}

func (u *Category) BeforeCreate(d *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
