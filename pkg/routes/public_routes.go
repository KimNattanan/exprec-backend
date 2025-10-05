package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"

	userHandler "github.com/KimNattanan/exprec-backend/internal/user/handler/rest"
	userRepository "github.com/KimNattanan/exprec-backend/internal/user/repository"
	userUseCase "github.com/KimNattanan/exprec-backend/internal/user/usecase"
)

func RegisterPublicRoutes(app fiber.Router, db *gorm.DB) {

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONTEND_URL"),
		AllowMethods: "GET,POST,PATCH,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	api := app.Group("/api/v2")

	// === Dependency Wiring ===

	userRepo := userRepository.NewGormUserRepository(db)
	userService := userUseCase.NewUserService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(userService, os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"))

	// === Public Routes ===

	authGroup := api.Group("/auth")
	authGroup.Get("/google/login", userHandler.GoogleLogin)
	authGroup.Get("/google/callback", userHandler.GoogleCallback)
}
