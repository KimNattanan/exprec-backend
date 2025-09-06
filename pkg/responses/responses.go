package responses

import (
	appError "github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
)

type MessageResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func Message(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(MessageResponse{Message: message})
}

func Error(c *fiber.Ctx, err error) error {
	return c.Status(appError.StatusCode(err)).JSON(ErrorResponse{Error: err.Error()})
}

func ErrorWithMessage(c *fiber.Ctx, err error, message string) error {
	return c.Status(appError.StatusCode(err)).JSON(ErrorResponse{Error: message})
}