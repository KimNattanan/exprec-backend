package dto

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID     uuid.UUID        `json:"id"`
	Email  string           `json:"email"`
	Name   string           `json:"name"`
	Prices []entities.Price `json:"prices"`
}
