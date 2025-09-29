package dto

type PreferencePatchRequest struct {
	Theme string `json:"theme" validate:"required"`
}
