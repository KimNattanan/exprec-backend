package middleware

import (
	appError "github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/KimNattanan/exprec-backend/pkg/responses"
	"github.com/KimNattanan/exprec-backend/pkg/token"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(secretKey string) fiber.Handler {

	tokenMaker := token.NewJWTMaker(secretKey)

	return func(c *fiber.Ctx) error {

		auth := c.Get("Authorization")
		if auth == "" {
			return responses.ErrorWithMessage(c, appError.ErrUnauthorized, "missing token")
		}
		tokenStr := auth[len("Bearer "):]

		claims, err := tokenMaker.VerfiyToken(tokenStr)
		if err != nil {
			return responses.ErrorWithMessage(c, appError.ErrUnauthorized, err.Error())
		}

		c.Locals("user_id", claims.ID)

		return c.Next()
	}
}
