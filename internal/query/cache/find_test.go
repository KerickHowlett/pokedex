package query_cache

import (
	"bytes"
	"testing"
	"time"

	f "test_tools/fixtures"
)

func TestQueryCache_Find(t *testing.T) {
	setup := func() *QueryCache {
		queryCache := NewQueryCache(24*time.Hour, f.FrozenTime)
		queryCache.entry[key] = cacheEntry{value: []byte(value)}

		return queryCache
	}

	t.Run("Cache Hit", func(t *testing.T) {
		t.Run("should return the value of the cache entry if it exists.", func(t *testing.T) {
			cache := setup()
			if cachedValue, _ := cache.Find(key); !bytes.Equal(cachedValue, []byte(value)) {
				t.Errorf("Expected value to be %v, but got %v", value, cachedValue)
			}
		})

		t.Run("should return TRUE if the cache entry exists.", func(t *testing.T) {
			cache := setup()
			if _, entryFound := cache.Find(key); !entryFound {
				t.Error("Expected cache entry to exist, but it does not.")
			}
		})
	})

	t.Run("Cache Miss", func(t *testing.T) {
		t.Run("should return nil if the cache entry does not exist.", func(t *testing.T) {
			cache := setup()
			if cachedValue, _ := cache.Find(invalid_key); cachedValue != nil {
				t.Errorf("Expected value to be nil, but got %v", cachedValue)
			}
		})

		t.Run("should return FALSE if the cache entry does not exist.", func(t *testing.T) {
			cache := setup()
			if _, entryFound := cache.Find(invalid_key); entryFound {
				t.Error("Expected cache entry to not exist, but it does.")
			}
		})
	})
}
