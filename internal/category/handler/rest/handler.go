package rest

import (
	"github.com/KimNattanan/exprec-backend/internal/category/dto"
	"github.com/KimNattanan/exprec-backend/internal/category/usecase"
	"github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/KimNattanan/exprec-backend/pkg/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HttpCategoryHandler struct {
	categoryUseCase usecase.CategoryUseCase
}

func NewHttpCategoryHandler(useCase usecase.CategoryUseCase) *HttpCategoryHandler {
	return &HttpCategoryHandler{categoryUseCase: useCase}
}

func (h *HttpCategoryHandler) Save(c *fiber.Ctx) error {
	req := new(dto.CategorySaveRequest)
	if err := c.BodyParser(req); err != nil {
		return responses.Error(c, apperror.ErrInvalidData)
	}
	category, err := dto.FromCategorySaveRequest(req)
	if err != nil {
		return responses.Error(c, apperror.ErrInvalidData)
	}
	userID := c.Locals("user_id").(uuid.UUID)
	category.UserID = userID
	if err := h.categoryUseCase.Save(c.Context(), category); err != nil {
		return responses.Error(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToCategoryResponse(category))
}

func (h *HttpCategoryHandler) Patch(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.Error(c, apperror.ErrInvalidData)
	}
	req := new(dto.CategoryPatchRequest)
	if err := c.BodyParser(req); err != nil {
		return responses.Error(c, apperror.ErrInvalidData)
	}
	category, err := dto.FromCategoryPatchRequest(req)
	if err != nil {
		return responses.Error(c, apperror.ErrInvalidData)
	}
	newCategory, err := h.categoryUseCase.Patch(c.Context(), id, category)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToCategoryResponse(newCategory))
}

func (h *HttpCategoryHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.Error(c, apperror.ErrInvalidData)
	}
	if err := h.categoryUseCase.Delete(id); err != nil {
		return responses.Error(c, err)
	}
	return responses.Message(c, fiber.StatusOK, "category deleted")
}

func (h *HttpCategoryHandler) FindByUserID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	categories, err := h.categoryUseCase.FindByUserID(userID)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToCategoryResponseList(categories))
}
