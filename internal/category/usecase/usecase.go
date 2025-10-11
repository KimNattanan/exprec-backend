package usecase

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/category/repository"
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/internal/transaction"
	"github.com/google/uuid"
)

type CategoryService struct {
	categoryRepo repository.CategoryRepository
	txManager    transaction.TransactionManager
}

func NewCategoryService(categoryRepo repository.CategoryRepository, txManager transaction.TransactionManager) CategoryUseCase {
	return &CategoryService{
		categoryRepo: categoryRepo,
		txManager:    txManager,
	}
}

func (s *CategoryService) Save(ctx context.Context, category *entities.Category) error {
	return s.txManager.Do(ctx, func(txCtx context.Context) error {
		if err := s.categoryRepo.Save(txCtx, category); err != nil {
			return err
		}
		if category.PrevID != nil {
			if err := s.categoryRepo.PatchNext(txCtx, *category.PrevID, &category.ID); err != nil {
				return err
			}
		}
		if category.NextID != nil {
			if err := s.categoryRepo.PatchPrev(txCtx, *category.NextID, &category.ID); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *CategoryService) FindByID(id uuid.UUID) (*entities.Category, error) {
	return s.categoryRepo.FindByID(id)
}

func (s *CategoryService) FindByUserID(userID uuid.UUID) ([]*entities.Category, error) {
	return s.categoryRepo.FindByUserID(userID)
}

func (s *CategoryService) Patch(ctx context.Context, id uuid.UUID, category *entities.Category) (*entities.Category, error) {
	err := s.txManager.Do(ctx, func(txCtx context.Context) error {
		categoryOld, err := s.categoryRepo.FindByID(id)
		if err != nil {
			return err
		}
		if categoryOld.PrevID != category.PrevID {
			if categoryOld.PrevID != nil {
				if err := s.categoryRepo.PatchNext(txCtx, *categoryOld.PrevID, categoryOld.NextID); err != nil {
					return err
				}
			}
			if category.PrevID != nil {
				if err := s.categoryRepo.PatchNext(txCtx, *category.PrevID, &id); err != nil {
					return err
				}
			}
		}
		if categoryOld.NextID != category.NextID {
			if categoryOld.NextID != nil {
				if err := s.categoryRepo.PatchPrev(txCtx, *categoryOld.NextID, categoryOld.PrevID); err != nil {
					return err
				}
			}
			if category.NextID != nil {
				if err := s.categoryRepo.PatchPrev(txCtx, *category.NextID, &id); err != nil {
					return err
				}
			}
		}
		return s.categoryRepo.PatchValue(txCtx, id, category)
	})
	if err != nil {
		return nil, err
	}
	return s.categoryRepo.FindByID(id)
}

func (s *CategoryService) Delete(id uuid.UUID) error {
	return s.categoryRepo.Delete(id)
}
