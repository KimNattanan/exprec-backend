package dto

import "github.com/KimNattanan/exprec-backend/internal/entities"

func ToPreferenceResponse(preference *entities.Preference) *PreferenceResponse {
	return &PreferenceResponse{
		Theme: preference.Theme,
	}
}

func FromPreferencePatchRequest(preference *PreferencePatchRequest) *entities.Preference {
	return &entities.Preference{
		Theme: preference.Theme,
	}
}
