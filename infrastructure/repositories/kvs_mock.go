package repositories

import (
	"challenge-scrapy/domain"
	"context"
	"encoding/json"
	"errors"
	"github.com/melisource/fury_go-toolkit-kvs/pkg/kvs"
)

type InMemoryMockKvsClient struct {
	mocks []kvs.Item
}

func NewInMemoryMockKvsClient() (*InMemoryMockKvsClient, *InMemoryMockKvsClient) {
	return &InMemoryMockKvsClient{
		mocks: []kvs.Item{},
	}, nil
}

func (m *InMemoryMockKvsClient) SaveFullItem(ctx context.Context, key string, fullItemArray *[]domain.FullItem) error {

	err := m.Set(ctx, key, fullItemArray)
	if err != nil {
		return err
	}

	return nil
}
func (m *InMemoryMockKvsClient) GetFullItem(ctx context.Context, key string) (*[]domain.FullItem, error) {
	item, err := m.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	fItem := make([]domain.FullItem, 0)

	convert := item.Value
	_ = json.Unmarshal(convert, &fItem)

	return &fItem, nil
}

func (m *InMemoryMockKvsClient) Set(_ context.Context, key string, value any) error {

	valueBytes, _ := json.Marshal(&value)

	m.mocks = append(m.mocks, kvs.Item{
		Key:   key,
		Value: valueBytes,
	})
	return nil
}

func (m *InMemoryMockKvsClient) Get(_ context.Context, key string) (kvs.Item, error) {
	item, err := m.GetItem(key)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (m *InMemoryMockKvsClient) GetItem(key string) (kvs.Item, error) {
	for _, item := range m.mocks {
		if item.Key == key {
			return item, nil
		}
	}
	return kvs.Item{}, errors.New("not found")
}

func (m *InMemoryMockKvsClient) Delete(_ context.Context, key string) (bool, error) {
	_, err := m.GetItem(key)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *InMemoryMockKvsClient) BulkSet(_ context.Context, _ []kvs.Item) (kvs.Bulk, error) {
	return kvs.Bulk{}, nil
}

func (m *InMemoryMockKvsClient) BulkGet(_ context.Context, _ []string) (kvs.Bulk, error) {
	return kvs.Bulk{}, nil
}

func (m *InMemoryMockKvsClient) BulkDelete(_ context.Context, _ []string) (kvs.Bulk, error) {
	return kvs.Bulk{}, nil
}

func (m *InMemoryMockKvsClient) BulkWrite(_ context.Context, _ []kvs.WriteOperation) error {
	return nil
}

func (m *InMemoryMockKvsClient) BatchSet(_ context.Context, _ []kvs.Item) error {
	return nil
}

func (m *InMemoryMockKvsClient) BatchGet(_ context.Context, _ []string) (map[string]kvs.Item, error) {
	return map[string]kvs.Item{}, nil
}

func (m *InMemoryMockKvsClient) BatchDelete(_ context.Context, _ []string) error {
	return nil
}

func (m *InMemoryMockKvsClient) BatchWrite(_ context.Context, _ []kvs.WriteOperation) error {
	return nil
}

func (m *InMemoryMockKvsClient) Shutdown() {}
