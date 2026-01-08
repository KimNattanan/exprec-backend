package httpserver

import (
	"log"

	"github.com/KimNattanan/exprec-backend/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func Start(app *fiber.App, cfg *config.Config) {
	log.Println("Starting REST server on port:", cfg.AppPort)
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatalf("REST server error: %v", err)
	}
}

func Shutdown(app *fiber.App) {
	log.Println("Shutting down REST server...")
	if err := app.Shutdown(); err != nil {
		log.Printf("Error shutting down REST server: %v", err)
	}
}
