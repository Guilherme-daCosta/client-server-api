package repository

import (
	"context"
	"time"

	"github.com/Guilherme-daCosta/client-server-api/server/model"
	"gorm.io/gorm"
)

type currencyQuoteRepository struct {
	db *gorm.DB
}

func NewCurrencyQuoteRepository(db *gorm.DB) model.CurrencyQuoteRepository {
	return &currencyQuoteRepository{
		db: db,
	}
}

func (r *currencyQuoteRepository) Save(exchange *model.CurrencyQuote) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Millisecond)
	defer cancel()

	return r.db.WithContext(ctx).Create(exchange).Error
}
