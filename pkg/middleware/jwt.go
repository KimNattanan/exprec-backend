package middleware

import (
	"os"

	appError "github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/KimNattanan/exprec-backend/pkg/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("loginToken")
	jwtSecretKey := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil || !token.Valid {
		return responses.Error(c, appError.ErrUnauthorized)
	}

	claim := token.Claims.(jwt.MapClaims)
	user_id := claim["user_id"]

	c.Locals("user_id", user_id)

	return c.Next()
}
