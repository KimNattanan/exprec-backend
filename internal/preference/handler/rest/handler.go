package handler

import (
	"github.com/KimNattanan/exprec-backend/internal/preference/dto"
	"github.com/KimNattanan/exprec-backend/internal/preference/usecase"
	appError "github.com/KimNattanan/exprec-backend/pkg/apperror"
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
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	req := new(dto.PreferencePatchRequest)
	if err := c.BodyParser(req); err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	preference, err := h.preferenceUseCase.Patch(id, dto.FromPreferencePatchRequest(req))
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPreferenceResponse(preference))
}

func (h *HttpPreferenceHandler) FindByUserID(c *fiber.Ctx) error {
	user_id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	preference, err := h.preferenceUseCase.FindByUserID(user_id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPreferenceResponse(preference))
}
