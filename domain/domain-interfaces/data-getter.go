package domain_interfaces

import (
	"challenge-scrapy/domain"
	"context"
)

type DataBusinessGetter interface {
	GetItems(items []string) ([]domain.Item, error)
	GetCategory(categoryID string) (domain.CategoryName, error)
	GetCurrency(currencyID string) (domain.CurrencyDescription, error)
	GetUserNickname(sellerID int) (domain.NickName, error)
	SetContext(ctx context.Context)
	SetAuth(authManager *domain.AuthManager)
}
