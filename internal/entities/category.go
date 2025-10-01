package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	UserID  uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	ID      uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	PrevID  *uuid.UUID `gorm:"type:uuid" json:"prev_id"`
	NextID  *uuid.UUID `gorm:"type:uuid" json:"next_id"`
	Title   string     `gorm:"type:varchar(50)" json:"title"`
	BgColor string     `gorm:"type:char(7)" json:"bg_color"`

	Prev *Category `gorm:"foreignKey:PrevID;references:ID"`
	Next *Category `gorm:"foreignKey:NextID;references:ID"`
}

func (u *Category) BeforeCreate(db *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

func (p *Category) BeforeDelete(db *gorm.DB) (err error) {
	if p.PrevID != nil {
		if err := db.Model(&Category{}).
			Where("id = ?", p.PrevID).
			Update("next_id", p.NextID).Error; err != nil {
			return err
		}
	}
	if p.NextID != nil {
		if err := db.Model(&Category{}).
			Where("id = ?", p.NextID).
			Update("prev_id", p.PrevID).Error; err != nil {
			return err
		}
	}
	return nil
}
