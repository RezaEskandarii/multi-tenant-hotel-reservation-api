package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Manager struct {
	Client *redis.Client
	Ctx    context.Context
}

func (c *Manager) Set(key string, value interface{}, expiration time.Duration) error {

	return c.Client.Set(c.Ctx, key, value, expiration).Err()
}

func (c *Manager) Get(key string) (string, error) {
	return c.Client.Get(c.Ctx, key).Result()
}
