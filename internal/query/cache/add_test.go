package query_cache

import (
	"testing"
	"time"
)

func TestQueryCache_Add(t *testing.T) {
	runQueryCacheAddTest := func() (queryCache *QueryCache, currentTime time.Time) {
		queryCache = NewQueryCache()

		currentTime = time.Date(2024, time.April, 6, 12, 0, 0, 0, time.UTC)
		queryCache.Add(key, []byte(value), currentTime)

		return queryCache, currentTime
	}

	t.Run("should add new cache entry.", func(t *testing.T) {
		cache, _ := runQueryCacheAddTest()
		if _, exists := cache.entry[key]; !exists {
			t.Errorf("Expected entry with key '%s' to exist in cache, but it doesn't", key)
		}
	})

	t.Run("should set createdAt as the current time.", func(t *testing.T) {
		cache, expectedTime := runQueryCacheAddTest()
		if entry := cache.entry[key]; !entry.createdAt.Equal(expectedTime) {
			t.Errorf("Expected createdAt to be %v, but got %v", expectedTime.String(), entry.createdAt.String())
		}
	})

	t.Run("should set value as the provided value with it unmutated.", func(t *testing.T) {
		cache, _ := runQueryCacheAddTest()
		if entry := cache.entry[key]; string(entry.value) != value {
			t.Errorf("Expected value to be '%s', but got '%s'", value, string(entry.value))
		}
	})
}
