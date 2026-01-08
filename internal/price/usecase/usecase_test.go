package usecase

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/pkg/transaction"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type mockPriceRepo struct {
	save         func(ctx context.Context, price *entities.Price) error
	findByID     func(id uuid.UUID) (*entities.Price, error)
	findByUserID func(userID uuid.UUID) ([]*entities.Price, error)
	patchValue   func(ctx context.Context, id uuid.UUID, price *entities.Price) error
	patchPrev    func(ctx context.Context, id uuid.UUID, prevID *uuid.UUID) error
	patchNext    func(ctx context.Context, id uuid.UUID, nextID *uuid.UUID) error
	delete       func(id uuid.UUID) error
}

func (m *mockPriceRepo) Save(ctx context.Context, price *entities.Price) error {
	return m.save(ctx, price)
}
func (m *mockPriceRepo) FindByID(id uuid.UUID) (*entities.Price, error) {
	return m.findByID(id)
}
func (m *mockPriceRepo) FindByUserID(userID uuid.UUID) ([]*entities.Price, error) {
	return m.findByUserID(userID)
}
func (m *mockPriceRepo) PatchValue(ctx context.Context, id uuid.UUID, price *entities.Price) error {
	return m.patchValue(ctx, id, price)
}
func (m *mockPriceRepo) PatchPrev(ctx context.Context, id uuid.UUID, prevID *uuid.UUID) error {
	return m.patchPrev(ctx, id, prevID)
}
func (m *mockPriceRepo) PatchNext(ctx context.Context, id uuid.UUID, nextID *uuid.UUID) error {
	return m.patchNext(ctx, id, nextID)
}
func (m *mockPriceRepo) Delete(id uuid.UUID) error {
	return m.delete(id)
}

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm database: %v", err)
	}

	t.Run("success; no L, R", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectCommit()

		uID := uuid.New()

		repo := &mockPriceRepo{
			save: func(ctx context.Context, price *entities.Price) error {
				price.ID = uID
				return nil
			},
			patchValue: func(ctx context.Context, id uuid.UUID, price *entities.Price) error {
				return gorm.ErrRecordNotFound
			},
			patchPrev: func(ctx context.Context, id uuid.UUID, prevID *uuid.UUID) error {
				return gorm.ErrRecordNotFound
			},
			patchNext: func(ctx context.Context, id uuid.UUID, nextID *uuid.UUID) error {
				return gorm.ErrRecordNotFound
			},
		}
		txManager := transaction.NewGormTxManager(gormDB)
		service := NewPriceService(repo, txManager)

		price := entities.Price{
			PrevID: nil,
			NextID: nil,
		}
		err := service.Save(context.TODO(), &price)

		require.NoError(t, err)
		assert.Equal(t, uID, price.ID)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("success; has L, no R", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectCommit()

		uID := uuid.New()
		l := entities.Price{
			ID:     uuid.New(),
			PrevID: nil,
			NextID: nil,
		}

		repo := &mockPriceRepo{
			save: func(ctx context.Context, price *entities.Price) error {
				price.ID = uID
				return nil
			},
			patchValue: func(ctx context.Context, id uuid.UUID, price *entities.Price) error {
				if id == l.ID {
					return nil
				}
				return gorm.ErrRecordNotFound
			},
			patchPrev: func(ctx context.Context, id uuid.UUID, prevID *uuid.UUID) error {
				if id == l.ID {
					l.PrevID = prevID
					return nil
				}
				return gorm.ErrRecordNotFound
			},
			patchNext: func(ctx context.Context, id uuid.UUID, nextID *uuid.UUID) error {
				if id == l.ID {
					l.NextID = nextID
					return nil
				}
				return gorm.ErrRecordNotFound
			},
		}
		txManager := transaction.NewGormTxManager(gormDB)
		service := NewPriceService(repo, txManager)

		price := entities.Price{
			PrevID: &l.ID,
			NextID: nil,
		}
		err := service.Save(context.TODO(), &price)

		require.NoError(t, err)
		assert.Equal(t, uID, price.ID, "uID incorrect")
		assert.Nil(t, l.PrevID, "L->L incorrect")
		assert.Equal(t, &uID, l.NextID, "L->R incorrect")

		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("success; no L, has R", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectCommit()

		uID := uuid.New()
		r := entities.Price{
			ID:     uuid.New(),
			PrevID: nil,
			NextID: nil,
		}

		repo := &mockPriceRepo{
			save: func(ctx context.Context, price *entities.Price) error {
				price.ID = uID
				return nil
			},
			patchValue: func(ctx context.Context, id uuid.UUID, price *entities.Price) error {
				if id == r.ID {
					return nil
				}
				return gorm.ErrRecordNotFound
			},
			patchPrev: func(ctx context.Context, id uuid.UUID, prevID *uuid.UUID) error {
				if id == r.ID {
					r.PrevID = prevID
					return nil
				}
				return gorm.ErrRecordNotFound
			},
			patchNext: func(ctx context.Context, id uuid.UUID, nextID *uuid.UUID) error {
				if id == r.ID {
					r.NextID = nextID
					return nil
				}
				return gorm.ErrRecordNotFound
			},
		}
		txManager := transaction.NewGormTxManager(gormDB)
		service := NewPriceService(repo, txManager)

		price := entities.Price{
			PrevID: nil,
			NextID: &r.ID,
		}
		err := service.Save(context.TODO(), &price)

		require.NoError(t, err)
		assert.Equal(t, uID, price.ID, "uID incorrect")
		assert.Equal(t, &uID, r.PrevID, "R->L incorrect")
		assert.Nil(t, r.NextID, "R->R incorrect")

		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("success; has L, R", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectCommit()

		uID := uuid.New()
		l := entities.Price{
			ID:     uuid.New(),
			PrevID: nil,
			NextID: nil,
		}
		r := entities.Price{
			ID:     uuid.New(),
			PrevID: nil,
			NextID: nil,
		}

		repo := &mockPriceRepo{
			save: func(ctx context.Context, price *entities.Price) error {
				price.ID = uID
				return nil
			},
			patchValue: func(ctx context.Context, id uuid.UUID, price *entities.Price) error {
				switch id {
				case l.ID:
					return nil
				case r.ID:
					return nil
				}
				return gorm.ErrRecordNotFound
			},
			patchPrev: func(ctx context.Context, id uuid.UUID, prevID *uuid.UUID) error {
				switch id {
				case l.ID:
					l.PrevID = prevID
					return nil
				case r.ID:
					r.PrevID = prevID
					return nil
				}
				return gorm.ErrRecordNotFound
			},
			patchNext: func(ctx context.Context, id uuid.UUID, nextID *uuid.UUID) error {
				switch id {
				case l.ID:
					l.NextID = nextID
					return nil
				case r.ID:
					r.NextID = nextID
					return nil
				}
				return gorm.ErrRecordNotFound
			},
		}
		txManager := transaction.NewGormTxManager(gormDB)
		service := NewPriceService(repo, txManager)

		price := entities.Price{
			PrevID: &l.ID,
			NextID: &r.ID,
		}
		err := service.Save(context.TODO(), &price)

		require.NoError(t, err)
		assert.Equal(t, uID, price.ID, "uID incorrect")
		assert.Nil(t, l.PrevID, "L->L incorrect")
		assert.Equal(t, &uID, l.NextID, "L->R incorrect")
		assert.Equal(t, &uID, r.PrevID, "R->L incorrect")
		assert.Nil(t, r.NextID, "R->R incorrect")

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
