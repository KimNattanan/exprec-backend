package dto

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserPatch struct {
	Name string `json:"name" validate:"required"`
}

type GoogleLoginRequest struct {
	Token string `json:"token"`
}

type GoogleRegisterRequest struct {
	Token    string `json:"token"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
}
