package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Record struct {
	UserID          uuid.UUID `gorm:"type:uuid;primaryKey;index:idx_user_record;priority:1" json:"user_id"`
	ID              uuid.UUID `gorm:"type:uuid" json:"id"`
	CreatedAt       time.Time `gorm:"type:timestamptz(3);primaryKey;index:idx_user_record;priority:2" json:"created_at"`
	Amount          float32   `json:"amount"`
	Category        string    `gorm:"type:varchar(50)" json:"category"`
	CategoryBgColor string    `gorm:"type:char(7)" json:"category_bg_color"`
	Note            string    `gorm:"type:varchar(255)" json:"note"`
}

func (u *Record) BeforeCreate(db *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
