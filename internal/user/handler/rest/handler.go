package handler

import (
	"github.com/KimNattanan/exprec-backend/internal/user/dto"
	"github.com/KimNattanan/exprec-backend/internal/user/usecase"
	appError "github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/KimNattanan/exprec-backend/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpUserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewHttpUserHandler(useCase usecase.UserUseCase) *HttpUserHandler {
	return &HttpUserHandler{userUseCase: useCase}
}

func (h *HttpUserHandler) Register(c *fiber.Ctx) error {
	req := new(dto.RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}

	user := dto.ToUserEntity(req)
	if err := h.userUseCase.Register(user); err != nil {
		return responses.Error(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToUserResponse(user))
}

func (h *HttpUserHandler) Login(c *fiber.Ctx) error {
	req := new(dto.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}

	token, user, err := h.userUseCase.Login(req.Email, req.Password)
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrUnauthorized, "invalid email or password")
	}

	return c.JSON(fiber.Map{
		"user":  dto.ToUserResponse(user),
		"token": token,
	})
}

func (h *HttpUserHandler) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return responses.Error(c, appError.ErrInvalidData)
	}

	// user, err := h.userUseCase.FindByID(id)

	return nil
}
