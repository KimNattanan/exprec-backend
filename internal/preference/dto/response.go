package dto

import "github.com/google/uuid"

type PreferenceResponse struct {
	UserID uuid.UUID `json:"user_id"`
	Theme  string    `json:"theme"`
}
