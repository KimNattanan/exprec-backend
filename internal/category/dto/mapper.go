package dto

import (
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/google/uuid"
)

func ToCategoryResponse(category *entities.Category) *CategoryResponse {
	var (
		prevID = ""
		nextID = ""
	)
	if category.PrevID != nil {
		prevID = category.PrevID.String()
	}
	if category.NextID != nil {
		nextID = category.NextID.String()
	}
	return &CategoryResponse{
		ID:      category.ID.String(),
		PrevID:  prevID,
		NextID:  nextID,
		Title:   category.Title,
		BgColor: category.BgColor,
	}
}

func ToCategoryResponseList(categories []*entities.Category) []*CategoryResponse {
	result := make([]*CategoryResponse, len(categories))
	for i, u := range categories {
		result[i] = ToCategoryResponse(u)
	}
	return result
}

func FromCategorySaveRequest(category *CategorySaveRequest) (*entities.Category, error) {
	var (
		err       error
		prevID    uuid.UUID
		prevIDPtr *uuid.UUID
		nextID    uuid.UUID
		nextIDPtr *uuid.UUID
	)
	if err != nil {
		return nil, err
	}
	if category.PrevID != "" {
		prevID, err = uuid.Parse(category.PrevID)
		if err != nil {
			return nil, err
		}
		prevIDPtr = &prevID
	}
	if category.NextID != "" {
		nextID, err = uuid.Parse(category.NextID)
		if err != nil {
			return nil, err
		}
		nextIDPtr = &nextID
	}
	return &entities.Category{
		PrevID:  prevIDPtr,
		NextID:  nextIDPtr,
		Title:   category.Title,
		BgColor: category.BgColor,
	}, nil
}

func FromCategoryPatchRequest(category *CategoryPatchRequest) (*entities.Category, error) {
	var (
		err       error
		prevID    uuid.UUID
		prevIDPtr *uuid.UUID
		nextID    uuid.UUID
		nextIDPtr *uuid.UUID
	)
	if category.PrevID != "" {
		prevID, err = uuid.Parse(category.PrevID)
		if err != nil {
			return nil, err
		}
		prevIDPtr = &prevID
	}
	if category.NextID != "" {
		nextID, err = uuid.Parse(category.NextID)
		if err != nil {
			return nil, err
		}
		nextIDPtr = &nextID
	}
	return &entities.Category{
		PrevID:  prevIDPtr,
		NextID:  nextIDPtr,
		Title:   category.Title,
		BgColor: category.BgColor,
	}, nil
}
