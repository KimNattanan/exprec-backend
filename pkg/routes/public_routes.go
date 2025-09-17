package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	userHandler "github.com/KimNattanan/exprec-backend/internal/user/handler/rest"
	userRepository "github.com/KimNattanan/exprec-backend/internal/user/repository"
	userUseCase "github.com/KimNattanan/exprec-backend/internal/user/usecase"
)

func RegisterPublicRoutes(app fiber.Router, db *gorm.DB) {

	api := app.Group("/api/v1")

	// === Dependency Wiring ===

	userRepo := userRepository.NewGormUserRepository(db)
	userService := userUseCase.NewUserService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(userService, os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"))

	// === Public Routes ===

	authGroup := api.Group("/auth")
	authGroup.Post("/signup", userHandler.Register)
	authGroup.Post("/signin", userHandler.Login)
	authGroup.Get("/google/login", userHandler.GoogleLogin)
	authGroup.Get("/google/callback", userHandler.GoogleCallback)

	userGroup := api.Group("/users")
	userGroup.Get("/:id", userHandler.FindUserByID)
	// userGroup.Get("/:id/prices", priceHandler.FindAllPricesByUserID)
}
