package app

import (
	"log"

	"github.com/KimNattanan/exprec-backend/pkg/database"
	"github.com/KimNattanan/exprec-backend/pkg/middleware"
	"github.com/KimNattanan/exprec-backend/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func setupDependencies(env string) (*gorm.DB, error) {
	envFile := ".env"
	if env != "" {
		envFile = ".env." + env
	}
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: could not load .env file: %v", err)
	}

	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func setupRestServer(db *gorm.DB) *fiber.App {
	app := fiber.New()
	middleware.FiberMiddleware(app)
	routes.RegisterPublicRoutes(app, db)
	routes.RegisterPrivateRoutes(app, db)
	routes.RegisterNotFoundRoute(app)
	return app
}
