package dto

type PatchPreferenceRequest struct {
	Theme string `json:"theme" validate:"required"`
}
