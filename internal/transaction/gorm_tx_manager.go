package transaction

import (
	"context"

	"gorm.io/gorm"
)

type GormTxManager struct {
	db *gorm.DB
}

func NewGormTxManager(db *gorm.DB) TransactionManager {
	return &GormTxManager{db: db}
}

func (m *GormTxManager) Do(ctx context.Context, fn func(txCtx context.Context) error) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// store tx in context so repositories can pick it up
		txCtx := context.WithValue(ctx, "tx", tx)
		return fn(txCtx)
	})
}
