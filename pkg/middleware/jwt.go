package middleware

import (
	"github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/KimNattanan/exprec-backend/pkg/responses"
	"github.com/KimNattanan/exprec-backend/pkg/token"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func JWTMiddleware(secretKey string) fiber.Handler {

	tokenMaker := token.NewJWTMaker(secretKey)

	return func(c *fiber.Ctx) error {

		auth := c.Get("Authorization")
		if auth == "" {
			return responses.ErrorWithMessage(c, apperror.ErrUnauthorized, "missing token")
		}
		tokenStr := auth[len("Bearer "):]

		claims, err := tokenMaker.VerfiyToken(tokenStr)
		if err != nil {
			return responses.ErrorWithMessage(c, apperror.ErrUnauthorized, "invalid token")
		}

		userID, err :=uuid.Parse(claims.ID)
		if err != nil {
			return responses.Error(c, apperror.ErrUnauthorized)
		}

		c.Locals("user_id", userID)

		return c.Next()
	}
}
