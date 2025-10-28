package database

import (
	"fmt"
	"os"
	"time"

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
		// isProd   = os.Getenv("ENV") == "production"
		sslmode = "disable"
	)
	// if isProd {
	// 	sslmode = "require"
	// }
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)
	fmt.Println("check DSN:", dsn)

	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Connected to database!")
			break
		}
		fmt.Printf("Database not ready, retrying in 2 seconds... (%d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}
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
