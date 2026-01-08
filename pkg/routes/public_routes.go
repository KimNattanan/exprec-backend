package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	sessionHandler "github.com/KimNattanan/exprec-backend/internal/session/handler/rest"
	sessionRepository "github.com/KimNattanan/exprec-backend/internal/session/repository"
	sessionUseCase "github.com/KimNattanan/exprec-backend/internal/session/usecase"
	"github.com/KimNattanan/exprec-backend/pkg/config"

	userHandler "github.com/KimNattanan/exprec-backend/internal/user/handler/rest"
	userRepository "github.com/KimNattanan/exprec-backend/internal/user/repository"
	userUseCase "github.com/KimNattanan/exprec-backend/internal/user/usecase"
)

func RegisterPublicRoutes(app fiber.Router, db *gorm.DB, cfg *config.Config) {
	api := app.Group("/api/v1")

	// === Dependency Wiring ===

	sessionRepo := sessionRepository.NewGormSessionRepository(db)
	sessionService := sessionUseCase.NewSessionService(sessionRepo)
	userRepo := userRepository.NewGormUserRepository(db)
	userService := userUseCase.NewUserService(userRepo, cfg.JWTSecret)
	userHandler := userHandler.NewHttpUserHandler(
		userService,
		cfg.GoogleClientID,
		cfg.GoogleClientSecret,
		cfg.GoogleRedirectURL,
		cfg.JWTSecret,
		sessionService,
		cfg.AppEnv,
		cfg.AppDomain,
		cfg.FrontendRedirectURL,
	)
	sessionHandler := sessionHandler.NewHttpSessionHandler(
		sessionService,
		userService,
		cfg.JWTSecret,
		cfg.AppDomain,
	)

	// === Public Routes ===

	authGroup := api.Group("/auth")
	authGroup.Get("/google/login", userHandler.GoogleLogin)
	authGroup.Get("/google/callback", userHandler.GoogleCallback)
	authGroup.Post("/refresh", sessionHandler.RenewToken)
	authGroup.Post("/logout", sessionHandler.Logout)
}
