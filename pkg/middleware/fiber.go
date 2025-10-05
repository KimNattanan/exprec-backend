package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddleware(app *fiber.App) {
	app.Use(
		logger.New(),
		cors.New(cors.Config{
			AllowOrigins:     os.Getenv("FRONTEND_URL"),
			AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
			AllowCredentials: true,
		}),
	)
}
