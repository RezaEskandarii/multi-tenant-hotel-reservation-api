package common_services

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type CacheManager interface {
	Set(key string, value interface{}, expiration *time.Duration) error
	Get(key string) (string, error)
	Del(key string) (int64, error)
	Update(key string, value interface{}) error
}

type CacheService struct {
	Client            *redis.Client
	Ctx               context.Context
	DefaultExpiration *time.Duration
}

func NewCacheService(addr, password string, db int, ctx context.Context) *CacheService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return &CacheService{
		Client:            rdb,
		Ctx:               ctx,
		DefaultExpiration: nil,
	}
}

// Set stores given value with key in cache storage.
func (m *CacheService) Set(key string, value interface{}, expiration *time.Duration) error {
	if expiration == nil {
		expiration = m.DefaultExpiration
	}
	return m.Client.Set(m.Ctx, key, value, *expiration).Err()
}

// Get returns cache value by given key.
func (m *CacheService) Get(key string) (string, error) {
	return m.Client.Get(m.Ctx, key).Result()
}

// Del remove cache by given key.
func (m *CacheService) Del(key string) (int64, error) {
	return m.Client.Del(m.Ctx, key).Result()
}

// Update removes old cache and puts new cache with given key.
func (m *CacheService) Update(key string, value interface{}) error {

	data, err := m.Client.Get(m.Ctx, key).Result()
	if err != nil {
		return err
	}
	if data != "" {
		m.Client.Del(m.Ctx, key)
	}

	return m.Set(key, value, m.DefaultExpiration)
}
