package usecase

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/internal/price/repository"
	"github.com/KimNattanan/exprec-backend/internal/transaction"
	"github.com/google/uuid"
)

type PriceService struct {
	priceRepo repository.PriceRepository
	txManager transaction.TransactionManager
}

func NewPriceService(priceRepo repository.PriceRepository, txManager transaction.TransactionManager) PriceUseCase {
	return &PriceService{
		priceRepo: priceRepo,
		txManager: txManager,
	}
}

func (s *PriceService) Save(ctx context.Context, price *entities.Price) error {
	return s.txManager.Do(ctx, func(txCtx context.Context) error {
		if err := s.priceRepo.Save(txCtx, price); err != nil {
			return err
		}
		if price.PrevID != nil {
			if err := s.priceRepo.PatchNext(txCtx, *price.PrevID, &price.ID); err != nil {
				return err
			}
		}
		if price.NextID != nil {
			if err := s.priceRepo.PatchPrev(txCtx, *price.NextID, &price.ID); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *PriceService) FindByID(id uuid.UUID) (*entities.Price, error) {
	return s.priceRepo.FindByID(id)
}

func (s *PriceService) FindByUserID(user_id uuid.UUID) ([]*entities.Price, error) {
	return s.priceRepo.FindByUserID(user_id)
}

func (s *PriceService) Patch(ctx context.Context, id uuid.UUID, price *entities.Price) (*entities.Price, error) {
	err := s.txManager.Do(ctx, func(txCtx context.Context) error {
		priceOld, err := s.priceRepo.FindByID(id)
		if err != nil {
			return err
		}
		if priceOld.PrevID != price.PrevID {
			if priceOld.PrevID != nil {
				if err := s.priceRepo.PatchNext(txCtx, *priceOld.PrevID, priceOld.NextID); err != nil {
					return err
				}
			}
			if price.PrevID != nil {
				if err := s.priceRepo.PatchNext(txCtx, *price.PrevID, &id); err != nil {
					return err
				}
			}
		}
		if priceOld.NextID != price.NextID {
			if priceOld.NextID != nil {
				if err := s.priceRepo.PatchPrev(txCtx, *priceOld.NextID, priceOld.PrevID); err != nil {
					return err
				}
			}
			if price.NextID != nil {
				if err := s.priceRepo.PatchPrev(txCtx, *price.NextID, &id); err != nil {
					return err
				}
			}
		}
		return s.priceRepo.PatchValue(txCtx, id, price)
	})
	if err != nil {
		return nil, err
	}
	return s.priceRepo.FindByID(id)
}

func (s *PriceService) Delete(id uuid.UUID) error {
	return s.priceRepo.Delete(id)
}
