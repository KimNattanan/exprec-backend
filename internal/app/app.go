package app

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/pkg/config"
	"github.com/KimNattanan/exprec-backend/pkg/database"
	"github.com/KimNattanan/exprec-backend/pkg/middleware"
	"github.com/KimNattanan/exprec-backend/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setupDependencies(env string) (*config.Config, *gorm.DB, error) {
	cfg := config.LoadConfig(env)

	db, err := database.Connect(cfg.DBDSN)
	if err != nil {
		return nil, nil, err
	}

	if env == "test" {
		db.Migrator().DropTable(
			&entities.User{},
			&entities.Preference{},
			&entities.Price{},
			&entities.Category{},
			&entities.Record{},
			&entities.Session{},
		)
	}
	if err := db.Migrator().AutoMigrate(
		&entities.User{},
		&entities.Preference{},
		&entities.Price{},
		&entities.Category{},
		&entities.Record{},
		&entities.Session{},
	); err != nil {
		return nil, nil, err
	}

	return cfg, db, nil
}

func setupRestServer(db *gorm.DB, cfg *config.Config) *fiber.App {
	app := fiber.New()
	middleware.FiberMiddleware(app, cfg.FrontendURL)
	routes.RegisterPublicRoutes(app, db, cfg)
	routes.RegisterPrivateRoutes(app, db, cfg)
	routes.RegisterNotFoundRoute(app)
	return app
}
