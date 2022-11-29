package common_services

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reservation-api/internal/global_variables"
	"testing"
)

func TestCanConnectToRedis(t *testing.T) {

	cfg := global_variables.Config{}

	testCases := []struct {
		key   string
		value interface{}
	}{
		{
			key:   "test_name",
			value: "reza",
		}, {
			key:   "test_foo",
			value: "Bar",
		}, {
			key:   "test_count",
			value: 10,
		}, {
			key:   "test_amount",
			value: 85.33,
		},
	}

	cacheService := NewCacheService(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.CacheDB, context.Background())

	for _, test := range testCases {

		t.Run("test_can_set_cache", func(t *testing.T) {
			err := cacheService.Set(test.key, test.value, nil)
			assert.NotNil(t, err)
		})

		t.Run("test_can_get_cache", func(t *testing.T) {
			value, err := cacheService.Get(test.key)
			assert.Nil(t, err)
			assert.Equal(t, value, test.value)
		})

		t.Run("test_can_update_cache", func(t *testing.T) {
			valueToUpdate := fmt.Sprintf("%s__updated", test.value)
			err := cacheService.Update(test.key, valueToUpdate)
			assert.Nil(t, err)
			updatedValue, err := cacheService.Get(test.key)
			assert.Equal(t, updatedValue, valueToUpdate)
		})

		t.Run("test_can_delete_cache", func(t *testing.T) {
			affectedRows, err := cacheService.Del(test.key)
			assert.Nil(t, err)
			assert.True(t, affectedRows > 0)
		})

		t.Run("test_value_is_empty_after_delete_cache", func(t *testing.T) {
			val, _ := cacheService.Get(test.key)
			assert.Nil(t, val)
			assert.Equal(t, val, "")
		})

	}

}
