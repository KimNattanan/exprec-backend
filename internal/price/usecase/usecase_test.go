package usecase

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KimNattanan/exprec-backend/internal/entities"
	"github.com/KimNattanan/exprec-backend/internal/transaction"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type mockPriceRepo struct {
	save         func(ctx context.Context, price *entities.Price) error
	findByID     func(id uuid.UUID) (*entities.Price, error)
	findByUserID func(user_id uuid.UUID) ([]*entities.Price, error)
	patch        func(ctx context.Context, id uuid.UUID, price *entities.Price) error
	delete       func(id uuid.UUID) error
}

func (m *mockPriceRepo) Save(ctx context.Context, price *entities.Price) error {
	return m.save(ctx, price)
}
func (m *mockPriceRepo) FindByID(id uuid.UUID) (*entities.Price, error) {
	return m.findByID(id)
}
func (m *mockPriceRepo) FindByUserID(user_id uuid.UUID) ([]*entities.Price, error) {
	return m.findByUserID(user_id)
}
func (m *mockPriceRepo) Patch(ctx context.Context, id uuid.UUID, price *entities.Price) error {
	return m.patch(ctx, id, price)
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
			patch: func(ctx context.Context, id uuid.UUID, price *entities.Price) error {
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
			patch: func(ctx context.Context, id uuid.UUID, price *entities.Price) error {
				if id == l.ID {
					l.PrevID = price.PrevID
					l.NextID = price.NextID
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
			patch: func(ctx context.Context, id uuid.UUID, price *entities.Price) error {
				if id == r.ID {
					r.PrevID = price.PrevID
					r.NextID = price.NextID
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
			patch: func(ctx context.Context, id uuid.UUID, price *entities.Price) error {
				switch id {
				case l.ID:
					l.PrevID = price.PrevID
					l.NextID = price.NextID
					return nil
				case r.ID:
					r.PrevID = price.PrevID
					r.NextID = price.NextID
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
