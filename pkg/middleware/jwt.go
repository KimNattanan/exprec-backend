package middleware

import (
	"fmt"
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
		fmt.Println("1: ", err)
		fmt.Println("2: ", token)
		fmt.Println("3: ", cookie)
		fmt.Println("4: ", jwtSecretKey)
		return responses.Error(c, appError.ErrUnauthorized)
	}

	claim := token.Claims.(jwt.MapClaims)
	user_id := claim["user_id"]

	c.Locals("user_id", user_id)

	return c.Next()
}

func JWTMiddleware(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if auth == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
	}
	tokenStr := auth[len("Bearer "):]

	// tokenStr := c.Cookies("token") // Assuming the token is stored in a cookie named "token"
	// if tokenStr == "" {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
	// }

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"]
	c.Locals("user_id", userID)

	return c.Next()
}
