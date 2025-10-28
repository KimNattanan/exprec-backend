package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	sessionHandler "github.com/KimNattanan/exprec-backend/internal/session/handler/rest"
	sessionRepository "github.com/KimNattanan/exprec-backend/internal/session/repository"
	sessionUseCase "github.com/KimNattanan/exprec-backend/internal/session/usecase"

	userHandler "github.com/KimNattanan/exprec-backend/internal/user/handler/rest"
	userRepository "github.com/KimNattanan/exprec-backend/internal/user/repository"
	userUseCase "github.com/KimNattanan/exprec-backend/internal/user/usecase"
)

func RegisterPublicRoutes(app fiber.Router, db *gorm.DB) {
	api := app.Group("/api/v2")

	// === Dependency Wiring ===

	sessionRepo := sessionRepository.NewGormSessionRepository(db)
	sessionService := sessionUseCase.NewSessionService(sessionRepo)
	userRepo := userRepository.NewGormUserRepository(db)
	userService := userUseCase.NewUserService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(
		userService,
		os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
		os.Getenv("JWT_SECRET"),
		sessionService,
	)
	sessionHandler := sessionHandler.NewHttpSessionHandler(
		sessionService,
		userService,
		os.Getenv("JWT_SECRET"),
	)

	// === Public Routes ===

	authGroup := api.Group("/auth")
	authGroup.Get("/google/login", userHandler.GoogleLogin)
	authGroup.Get("/google/callback", userHandler.GoogleCallback)
	authGroup.Post("/refresh", sessionHandler.RenewToken)
	authGroup.Post("/logout", sessionHandler.Logout)
}
