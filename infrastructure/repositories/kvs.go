package repositories

import (
	"challenge-scrapy/domain"
	"context"
	"encoding/json"
	"github.com/melisource/fury_go-toolkit-kvs/pkg/kvs"
	"time"
)

const KvsReadTimeout = 400
const KvsWriteTimeout = 400

type KvsChallengeRepository struct {
	kvsClient kvs.Client
}

func NewKvsChallengeRepository(container string) (*KvsChallengeRepository, error) {
	optsFunc := []kvs.OptionClient{
		kvs.WithReadTimeout(KvsReadTimeout * time.Millisecond),
		kvs.WithWriteTimeout(KvsWriteTimeout * time.Millisecond),
	}

	kvsClient, err := kvs.NewClient(container, optsFunc...)
	if err != nil {
		return nil, err
	}

	return &KvsChallengeRepository{
		kvsClient: kvsClient,
	}, nil
}

func (k KvsChallengeRepository) GetFullItem(ctx context.Context, key string) ([]domain.FullItem, error) {

	arrayFullItem := make([]domain.FullItem, 0)
	fullItem, err := k.kvsClient.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fullItem.Value, &arrayFullItem)
	if err != nil {
		return nil, err
	}

	return arrayFullItem, nil
}

func (k KvsChallengeRepository) Save(ctx context.Context, key string, fullItemArray *[]domain.FullItem) error {

	err := k.kvsClient.Set(ctx, key, fullItemArray)
	if err != nil {
		return err
	}

	return nil
}
