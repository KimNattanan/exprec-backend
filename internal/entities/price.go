package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Price struct {
	UserID  uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	ID      uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	PrevID  *uuid.UUID `gorm:"type:uuid" json:"prev_id"`
	NextID  *uuid.UUID `gorm:"type:uuid" json:"next_id"`
	Amount  float32    `json:"amount"`
	BgColor string     `gorm:"type:char(7)" json:"bg_color"`

	Prev *Price `gorm:"foreignKey:PrevID;references:ID"`
	Next *Price `gorm:"foreignKey:NextID;references:ID"`
}

func (u *Price) BeforeCreate(d *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

func (p *Price) BeforeDelete(d *gorm.DB) (err error) {
	if p.PrevID != nil {
		if err := d.Model(&Price{}).
			Where("id = ?", p.PrevID).
			Update("next_id", p.NextID).Error; err != nil {
			return err
		}
	}
	if p.NextID != nil {
		if err := d.Model(&Price{}).
			Where("id = ?", p.NextID).
			Update("prev_id", p.PrevID).Error; err != nil {
			return err
		}
	}
	return nil
}
