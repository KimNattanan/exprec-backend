package handler

import (
	"github.com/KimNattanan/exprec-backend/internal/preference/usecase"
)

type HttpPreferenceHandler struct {
	userUseCase usecase.PreferenceUseCase
}

func NewHttpUserHandler(useCase usecase.PreferenceUseCase) *HttpPreferenceHandler {
	return &HttpPreferenceHandler{userUseCase: useCase}
}
