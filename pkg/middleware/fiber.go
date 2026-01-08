package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddleware(app *fiber.App, frontendURL string) {
	app.Use(
		logger.New(),
		cors.New(cors.Config{
			AllowOrigins:     frontendURL,
			AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
			AllowCredentials: true,
		}),
	)
}
