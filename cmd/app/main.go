package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	userHandler "github.com/KimNattanan/exprec-backend/internal/user/handler/rest"
	userRepository "github.com/KimNattanan/exprec-backend/internal/user/repository"
	userUseCase "github.com/KimNattanan/exprec-backend/internal/user/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: could not load .env file: %v", err)
	}

	app := fiber.New()

	const (
		host     = "localhost"
		port     = "5432"
		user     = "myuser"
		password = "mypassword"
		dbname   = "mydb"
	)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	if err := db.AutoMigrate(&entities.Preference{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	if err := db.AutoMigrate(&entities.Price{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	if err := db.AutoMigrate(&entities.Category{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	if err := db.AutoMigrate(&entities.Record{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	userRepo := userRepository.NewGormUserRepository(db)
	userService := userUseCase.NewUserService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(userService, os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	authGroup := app.Group("/auth")
	authGroup.Post("/signup", userHandler.Register)
	authGroup.Post("/signin", userHandler.Login)
	authGroup.Get("/google/login", userHandler.GoogleLogin)
	authGroup.Get("/google/callback", userHandler.GoogleCallback)

	app.Listen(":8000")
}
