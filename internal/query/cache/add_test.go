package query_cache

import (
	"testing"

	f "test_tools/fixtures"
)

func TestQueryCache_Add(t *testing.T) {
	runQueryCacheAddTest := func() (queryCache *QueryCache) {
		queryCache = NewQueryCache(mockedTTL)

		queryCache.Add(key, []byte(value), f.FrozenTime)

		return queryCache
	}

	t.Run("should add new cache entry.", func(t *testing.T) {
		cache := runQueryCacheAddTest()
		if _, exists := cache.entry[key]; !exists {
			t.Errorf("Expected entry with key '%s' to exist in cache, but it doesn't", key)
		}
	})

	t.Run("should set createdAt as the current time.", func(t *testing.T) {
		cache := runQueryCacheAddTest()
		if entry := cache.entry[key]; !entry.createdAt.Equal(f.FrozenTime) {
			t.Errorf("Expected createdAt to be %v, but got %v", f.FrozenTime.String(), entry.createdAt.String())
		}
	})

	t.Run("should set value as the provided value with it unmutated.", func(t *testing.T) {
		cache := runQueryCacheAddTest()
		if entry := cache.entry[key]; string(entry.value) != value {
			t.Errorf("Expected value to be '%s', but got '%s'", value, string(entry.value))
		}
	})
}
