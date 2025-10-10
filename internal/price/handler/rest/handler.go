package rest

import (
	"fmt"

	"github.com/KimNattanan/exprec-backend/internal/price/dto"
	"github.com/KimNattanan/exprec-backend/internal/price/usecase"
	appError "github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/KimNattanan/exprec-backend/pkg/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HttpPriceHandler struct {
	priceUseCase usecase.PriceUseCase
}

func NewHttpPriceHandler(useCase usecase.PriceUseCase) *HttpPriceHandler {
	return &HttpPriceHandler{priceUseCase: useCase}
}

func (h *HttpPriceHandler) Save(c *fiber.Ctx) error {
	req := new(dto.PriceSaveRequest)
	if err := c.BodyParser(req); err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	price, err := dto.FromPriceSaveRequest(req)
	if err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	userID, err := uuid.Parse(c.Locals("user_id").(string))
	if err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	price.UserID = userID
	if err := h.priceUseCase.Save(c.Context(), price); err != nil {
		return responses.Error(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToPriceResponse(price))
}

func (h *HttpPriceHandler) Patch(c *fiber.Ctx) error {
	id := c.Params("id")
	req := new(dto.PricePatchRequest)
	if err := c.BodyParser(req); err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	price, err := dto.FromPricePatchRequest(req)
	if err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	newPrice, err := h.priceUseCase.Patch(c.Context(), id, price)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPriceResponse(newPrice))
}

func (h *HttpPriceHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.priceUseCase.Delete(id); err != nil {
		return responses.Error(c, err)
	}
	return responses.Message(c, fiber.StatusOK, "price deleted")
}

func (h *HttpPriceHandler) FindByUserID(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	prices, err := h.priceUseCase.FindByUserID(fmt.Sprint(userID))
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPriceResponseList(prices))
}
