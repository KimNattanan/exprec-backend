package handler

import (
	"github.com/KimNattanan/exprec-backend/internal/preference/dto"
	"github.com/KimNattanan/exprec-backend/internal/preference/usecase"
	"github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/KimNattanan/exprec-backend/pkg/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HttpPreferenceHandler struct {
	preferenceUseCase usecase.PreferenceUseCase
}

func NewHttpPreferenceHandler(useCase usecase.PreferenceUseCase) *HttpPreferenceHandler {
	return &HttpPreferenceHandler{preferenceUseCase: useCase}
}

func (h *HttpPreferenceHandler) Patch(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	req := new(dto.PreferencePatchRequest)
	if err := c.BodyParser(req); err != nil {
		return responses.Error(c, apperror.ErrInvalidData)
	}
	preference, err := h.preferenceUseCase.Patch(userID, dto.FromPreferencePatchRequest(req))
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPreferenceResponse(preference))
}

func (h *HttpPreferenceHandler) FindByUserID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	preference, err := h.preferenceUseCase.FindByUserID(userID)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPreferenceResponse(preference))
}
