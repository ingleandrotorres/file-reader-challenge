package cache

import (
	"context"
	"github.com/mercadolibre/fury_go-toolkit-cache/pkg/cache"
	"sync"
	"time"
)

type InMemoryCache struct {
	data sync.Map
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{}
}

func (m *InMemoryCache) Shutdown() {}

func (m *InMemoryCache) Get(_ context.Context, key string) ([]byte, error) {
	value, ok := m.data.Load(key)
	if !ok {
		return nil, nil
	}
	return value.([]byte), nil
}

func (m *InMemoryCache) BulkGet(_ context.Context, keys ...string) (cache.BulkResponse, error) {
	response := make(cache.BulkResponse)
	for _, key := range keys {
		value, ok := m.data.Load(key)
		if ok {
			response[key] = value.([]byte)
		}
	}
	return response, nil
}

func (m *InMemoryCache) Set(_ context.Context, key string, value []byte, expiration time.Duration) error {
	m.data.Store(key, value)
	return nil
}

func (m *InMemoryCache) Delete(_ context.Context, keys ...string) error {
	for _, key := range keys {
		m.data.Delete(key)
	}
	return nil
}

func (m *InMemoryCache) Touch(_ context.Context, _ string, _ time.Duration) error {
	return nil
}
