package main

import (
	"fmt"
	"log"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	userHandler "github.com/KimNattanan/exprec-backend/internal/user/handler/rest"
	userRepository "github.com/KimNattanan/exprec-backend/internal/user/repository"
	userUseCase "github.com/KimNattanan/exprec-backend/internal/user/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
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
	userHandler := userHandler.NewHttpUserHandler(userService)

	app.Post("/auth/signup", userHandler.Register)
	app.Post("/auth/signin", userHandler.Login)

	app.Listen(":8000")
}
