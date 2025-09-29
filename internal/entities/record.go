package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Record struct {
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Amount    float32   `json:"amount"`
	Category  string    `gorm:"type:varchar(50)" json:"category"`
	Note      string    `gorm:"type:varchar(255)" json:"note"`
}

func (u *Record) BeforeCreate(d *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	return
}
