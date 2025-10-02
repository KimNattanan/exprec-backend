package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"

	"github.com/KimNattanan/exprec-backend/internal/transaction"

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

func RegisterPublicRoutes(app fiber.Router, db *gorm.DB) {

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // or specific domains
		AllowMethods: "GET,POST,PATCH,DELETE",
	}))

	api := app.Group("/api/v2")

	// === Dependency Wiring ===

	txManager := transaction.NewGormTxManager(db)

	userRepo := userRepository.NewGormUserRepository(db)
	userService := userUseCase.NewUserService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(userService, os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"))

	priceRepo := priceRepository.NewGormPriceRepository(db)
	priceService := priceUseCase.NewPriceService(priceRepo, txManager)
	priceHandler := priceHandler.NewHttpPriceHandler(priceService)

	categoryRepo := categoryRepository.NewGormCategoryRepository(db)
	categoryService := categoryUseCase.NewCategoryService(categoryRepo, txManager)
	categoryHandler := categoryHandler.NewHttpCategoryHandler(categoryService)

	recordRepo := recordRepository.NewGormRecordRepository(db)
	recordService := recordUseCase.NewRecordService(recordRepo)
	recordHandler := recordHandler.NewHttpRecordHandler(recordService)

	preferenceRepo := preferenceRepository.NewGormPreferenceRepository(db)
	preferenceService := preferenceUseCase.NewPreferenceService(preferenceRepo)
	preferenceHandler := preferenceHandler.NewHttpPreferenceHandler(preferenceService)

	// === Public Routes ===

	authGroup := api.Group("/auth")
	authGroup.Post("/signup", userHandler.Register)
	authGroup.Post("/signin", userHandler.Login)
	authGroup.Get("/google/login", userHandler.GoogleLogin)
	authGroup.Get("/google/callback", userHandler.GoogleCallback)

	userGroup := api.Group("/users")
	userGroup.Get("/:id", userHandler.FindUserByID)
	userGroup.Delete("/:id", userHandler.Delete)

	priceGroup := api.Group("/prices")
	priceGroup.Post("/", priceHandler.Save)
	priceGroup.Patch("/:id", priceHandler.Patch)
	priceGroup.Delete("/:id", priceHandler.Delete)
	priceGroup.Get("/user/:id", priceHandler.FindByUserID)

	categoryGroup := api.Group("/categories")
	categoryGroup.Post("/", categoryHandler.Save)
	categoryGroup.Patch("/:id", categoryHandler.Patch)
	categoryGroup.Delete("/:id", categoryHandler.Delete)
	categoryGroup.Get("/user/:id", categoryHandler.FindByUserID)

	recordGroup := api.Group("/records")
	recordGroup.Post("/", recordHandler.Save)
	recordGroup.Delete("/:id", recordHandler.Delete)
	recordGroup.Get("/user/:id", recordHandler.FindByUserID)

	preferenceGroup := api.Group("/preferences")
	preferenceGroup.Patch("/:id", preferenceHandler.Patch)
	preferenceGroup.Get("/:id", preferenceHandler.FindByUserID)
}
