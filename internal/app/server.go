package app

import (
	"fmt"
	"log"
	"os"

	"github.com/KimNattanan/exprec-backend/pkg/middleware"
	"github.com/KimNattanan/exprec-backend/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: could not load .env file: %v", err)
	}

	app := fiber.New()

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	middleware.FiberMiddleware(app)
	routes.RegisterPublicRoutes(app, db)
	routes.RegisterPrivateRoutes(app, db)

	app.Listen(":8000")
}
