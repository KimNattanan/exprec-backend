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
			if err := s.priceRepo.Patch(txCtx, *price.PrevID, &entities.Price{NextID: &price.ID}); err != nil {
				return err
			}
		}
		if price.NextID != nil {
			if err := s.priceRepo.Patch(txCtx, *price.NextID, &entities.Price{PrevID: &price.ID}); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *PriceService) FindByID(id uuid.UUID) (*entities.Price, error) {
	return s.FindByID(id)
}

func (s *PriceService) Patch(ctx context.Context, id uuid.UUID, price *entities.Price) (*entities.Price, error) {
	err := s.txManager.Do(ctx, func(txCtx context.Context) error {
		priceOld, err := s.priceRepo.FindByID(id)
		if err != nil {
			return err
		}
		if priceOld.PrevID != price.PrevID {
			if priceOld.PrevID != nil {
				if err := s.priceRepo.Patch(txCtx, *priceOld.PrevID, &entities.Price{NextID: priceOld.NextID}); err != nil {
					return err
				}
			}
			if price.PrevID != nil {
				if err := s.priceRepo.Patch(txCtx, *price.PrevID, &entities.Price{NextID: &id}); err != nil {
					return err
				}
			}
		}
		if priceOld.NextID != price.NextID {
			if priceOld.NextID != nil {
				if err := s.priceRepo.Patch(txCtx, *priceOld.NextID, &entities.Price{PrevID: priceOld.PrevID}); err != nil {
					return err
				}
			}
			if price.NextID != nil {
				if err := s.priceRepo.Patch(txCtx, *price.NextID, &entities.Price{PrevID: &id}); err != nil {
					return err
				}
			}
		}
		return s.priceRepo.Patch(txCtx, id, price)
	})
	if err != nil {
		return nil, err
	}
	return s.priceRepo.FindByID(id)
}

func (s *PriceService) Delete(id uuid.UUID) error {
	return s.priceRepo.Delete(id)
}
