package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/KimNattanan/exprec-backend/internal/transaction"
	"github.com/KimNattanan/exprec-backend/pkg/middleware"

	userHandler "github.com/KimNattanan/exprec-backend/internal/user/handler/rest"
	userRepository "github.com/KimNattanan/exprec-backend/internal/user/repository"
	userUseCase "github.com/KimNattanan/exprec-backend/internal/user/usecase"

	priceHandler "github.com/KimNattanan/exprec-backend/internal/price/handler/rest"
	priceRepository "github.com/KimNattanan/exprec-backend/internal/price/repository"
	priceUseCase "github.com/KimNattanan/exprec-backend/internal/price/usecase"

	categoryHandler "github.com/KimNattanan/exprec-backend/internal/category/handler/rest"
	categoryRepository "github.com/KimNattanan/exprec-backend/internal/category/repository"
	categoryUseCase "github.com/KimNattanan/exprec-backend/internal/category/usecase"

	recordHandler "github.com/KimNattanan/exprec-backend/internal/record/handler/rest"
	recordRepository "github.com/KimNattanan/exprec-backend/internal/record/repository"
	recordUseCase "github.com/KimNattanan/exprec-backend/internal/record/usecase"

	preferenceHandler "github.com/KimNattanan/exprec-backend/internal/preference/handler/rest"
	preferenceRepository "github.com/KimNattanan/exprec-backend/internal/preference/repository"
	preferenceUseCase "github.com/KimNattanan/exprec-backend/internal/preference/usecase"
)

func RegisterPrivateRoutes(app fiber.Router, db *gorm.DB) {
	api := app.Group("/api/v2", middleware.AuthRequired)

	// === Dependency Wiring ===

	txManager := transaction.NewGormTxManager(db)

	userRepo := userRepository.NewGormUserRepository(db)
	userService := userUseCase.NewUserService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(userService, os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"))

	preferenceRepo := preferenceRepository.NewGormPreferenceRepository(db)
	preferenceService := preferenceUseCase.NewPreferenceService(preferenceRepo)
	preferenceHandler := preferenceHandler.NewHttpPreferenceHandler(preferenceService)

	priceRepo := priceRepository.NewGormPriceRepository(db)
	priceService := priceUseCase.NewPriceService(priceRepo, txManager)
	priceHandler := priceHandler.NewHttpPriceHandler(priceService)

	categoryRepo := categoryRepository.NewGormCategoryRepository(db)
	categoryService := categoryUseCase.NewCategoryService(categoryRepo, txManager)
	categoryHandler := categoryHandler.NewHttpCategoryHandler(categoryService)

	recordRepo := recordRepository.NewGormRecordRepository(db)
	recordService := recordUseCase.NewRecordService(recordRepo)
	recordHandler := recordHandler.NewHttpRecordHandler(recordService)

	// === Public Routes ===

	api.Get("/me", userHandler.GetUser)
	api.Get("/me/logout", userHandler.Logout)

	userGroup := api.Group("/users")
	userGroup.Delete("/", userHandler.Delete)

	preferenceGroup := api.Group("/preferences")
	preferenceGroup.Patch("/", preferenceHandler.Patch)
	preferenceGroup.Get("/", preferenceHandler.FindByUserID)

	priceGroup := api.Group("/prices")
	priceGroup.Post("/", priceHandler.Save)
	priceGroup.Patch("/:id", priceHandler.Patch)
	priceGroup.Delete("/:id", priceHandler.Delete)
	priceGroup.Get("/", priceHandler.FindByUserID)

	categoryGroup := api.Group("/categories")
	categoryGroup.Post("/", categoryHandler.Save)
	categoryGroup.Patch("/:id", categoryHandler.Patch)
	categoryGroup.Delete("/:id", categoryHandler.Delete)
	categoryGroup.Get("/", categoryHandler.FindByUserID)

	recordGroup := api.Group("/records")
	recordGroup.Post("/", recordHandler.Save)
	recordGroup.Delete("/:id", recordHandler.Delete)
	recordGroup.Get("/", recordHandler.FindByUserID)
	recordGroup.Get("/dashboard-data", recordHandler.GetUserDashboardData)

}
