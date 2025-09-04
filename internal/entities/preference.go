package entities

import (
	"github.com/google/uuid"
)

type Preference struct {
	UserID uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	Theme  string    `gorm:"type:varchar(50)"`

	User User `gorm:"foreignKey:UserID;reference:ID"`
}
