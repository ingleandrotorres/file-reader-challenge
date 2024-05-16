package domain_interfaces

import (
	"challenge-scrapy/domain"
	"context"
)

type ChallengeRepository interface {
	SaveFullItem(ctx context.Context, key string, fullItemArray *[]domain.FullItem) error
	GetFullItem(ctx context.Context, key string) (*[]domain.FullItem, error)
}
