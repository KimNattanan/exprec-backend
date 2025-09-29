package dto

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

func ToPriceResponse(price *entities.Price) *PriceResponse {
	var (
		prevID = ""
		nextID = ""
	)
	if price.PrevID != nil {
		prevID = price.PrevID.String()
	}
	if price.NextID != nil {
		nextID = price.NextID.String()
	}
	return &PriceResponse{
		UserID:  price.UserID.String(),
		ID:      price.ID.String(),
		PrevID:  prevID,
		NextID:  nextID,
		Amount:  price.Amount,
		BgColor: price.BgColor,
	}
}

func ToPriceResponseList(prices []*entities.Price) []*PriceResponse {
	result := make([]*PriceResponse, len(prices))
	for i, u := range prices {
		result[i] = ToPriceResponse(u)
	}
	return result
}

func FromPriceSaveRequest(price *PriceSaveRequest) (*entities.Price, error) {
	var (
		err       error
		userID    uuid.UUID
		prevID    uuid.UUID
		prevIDPtr *uuid.UUID
		nextID    uuid.UUID
		nextIDPtr *uuid.UUID
	)
	userID, err = uuid.Parse(price.UserID)
	if err != nil {
		return nil, err
	}
	if price.PrevID != "" {
		prevID, err = uuid.Parse(price.PrevID)
		if err != nil {
			return nil, err
		}
		prevIDPtr = &prevID
	}
	if price.NextID != "" {
		nextID, err = uuid.Parse(price.NextID)
		if err != nil {
			return nil, err
		}
		nextIDPtr = &nextID
	}
	return &entities.Price{
		UserID:  userID,
		PrevID:  prevIDPtr,
		NextID:  nextIDPtr,
		Amount:  price.Amount,
		BgColor: price.BgColor,
	}, nil
}

func FromPricePatchRequest(price *PricePatchRequest) (*entities.Price, error) {
	var (
		err       error
		prevID    uuid.UUID
		prevIDPtr *uuid.UUID
		nextID    uuid.UUID
		nextIDPtr *uuid.UUID
	)
	if price.PrevID != "" {
		prevID, err = uuid.Parse(price.PrevID)
		if err != nil {
			return nil, err
		}
		prevIDPtr = &prevID
	}
	if price.NextID != "" {
		nextID, err = uuid.Parse(price.NextID)
		if err != nil {
			return nil, err
		}
		nextIDPtr = &nextID
	}
	return &entities.Price{
		PrevID:  prevIDPtr,
		NextID:  nextIDPtr,
		Amount:  price.Amount,
		BgColor: price.BgColor,
	}, nil
}
