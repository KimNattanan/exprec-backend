package usecase

import (
	"context"

	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/internal/price/repository"
	"github.com/KimNattanan/exprec-backend/internal/transaction"
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
			if err := s.priceRepo.PatchNext(txCtx, price.PrevID.String(), price.ID.String()); err != nil {
				return err
			}
		}
		if price.NextID != nil {
			if err := s.priceRepo.PatchPrev(txCtx, price.NextID.String(), price.ID.String()); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *PriceService) FindByID(id string) (*entities.Price, error) {
	return s.priceRepo.FindByID(id)
}

func (s *PriceService) FindByUserID(userID string) ([]*entities.Price, error) {
	return s.priceRepo.FindByUserID(userID)
}

func (s *PriceService) Patch(ctx context.Context, id string, price *entities.Price) (*entities.Price, error) {
	err := s.txManager.Do(ctx, func(txCtx context.Context) error {
		priceOld, err := s.priceRepo.FindByID(id)
		if err != nil {
			return err
		}
		if priceOld.PrevID != price.PrevID {
			if priceOld.PrevID != nil {
				if err := s.priceRepo.PatchNext(txCtx, priceOld.PrevID.String(), priceOld.NextID.String()); err != nil {
					return err
				}
			}
			if price.PrevID != nil {
				if err := s.priceRepo.PatchNext(txCtx, price.PrevID.String(), id); err != nil {
					return err
				}
			}
		}
		if priceOld.NextID != price.NextID {
			if priceOld.NextID != nil {
				if err := s.priceRepo.PatchPrev(txCtx, priceOld.NextID.String(), priceOld.PrevID.String()); err != nil {
					return err
				}
			}
			if price.NextID != nil {
				if err := s.priceRepo.PatchPrev(txCtx, price.NextID.String(), id); err != nil {
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

func (s *PriceService) Delete(id string) error {
	return s.priceRepo.Delete(id)
}
