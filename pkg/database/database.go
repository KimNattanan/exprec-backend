package database

import (
	"fmt"
	"os"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
		isProd   = os.Getenv("ENV") == "production"
		sslmode  = "disable"
	)
	if isProd {
		sslmode = "require"
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.Migrator().AutoMigrate(
		&entities.User{},
		&entities.Preference{},
		&entities.Price{},
		&entities.Category{},
		&entities.Record{},
		&entities.Session{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
