package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Manager struct {
	Client            *redis.Client
	Ctx               context.Context
	DefaultExpiration *time.Duration
}

// Set stores given value with key in cache storage.
func (m *Manager) Set(key string, value interface{}, expiration *time.Duration) error {
	if expiration == nil {
		expiration = m.DefaultExpiration
	}
	return m.Client.Set(m.Ctx, key, value, *expiration).Err()
}

// Get returns cache value by given key.
func (m *Manager) Get(key string) (string, error) {
	return m.Client.Get(m.Ctx, key).Result()
}

// Del remove cache by given key.
func (m *Manager) Del(key string) (int64, error) {
	return m.Client.Del(m.Ctx, key).Result()
}

// Update removes old cache and puts new cache with given key.
func (m *Manager) Update(key string, value interface{}) error {

	data, err := m.Client.Get(m.Ctx, key).Result()
	if err != nil {
		return err
	}
	if data != "" {
		m.Client.Del(m.Ctx, key)
	}
	m.Set(key, value, m.DefaultExpiration)
	return nil
}
