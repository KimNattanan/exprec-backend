package repository

import "github.com/KimNattanan/exprec-backend/internal/entities"

type PriceRepository interface {
	Save(price *entities.Price) error
}
