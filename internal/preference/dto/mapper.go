package dto

import "github.com/KimNattanan/exprec-backend/internal/entities"

func ToPreferenceResponse(preference *entities.Preference) *PreferenceResponse {
	return &PreferenceResponse{
		UserID: preference.UserID,
		Theme:  preference.Theme,
	}
}
