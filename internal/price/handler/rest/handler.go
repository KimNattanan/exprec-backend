package rest

import (
	"github.com/KimNattanan/exprec-backend/internal/price/usecase"
)

type HttpPriceHandler struct {
	priceUseCase usecase.PriceUseCase
}

func NewHttpPriceHandler(useCase usecase.PriceUseCase) *HttpPriceHandler {
	return &HttpPriceHandler{priceUseCase: useCase}
}
