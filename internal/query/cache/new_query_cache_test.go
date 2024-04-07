package query_cache

import (
	"testing"
	"time"

	f "test_tools/fixtures"
)

func TestNewQueryCache(t *testing.T) {
	expectEmptyCache := func(cache *QueryCache) {
		if len(cache.entry) != 0 {
			t.Error("Expected cache to be empty, but found one or more entries")
		}
	}
	waitForEvictionLoop := func() { time.Sleep(5 * time.Millisecond) }

	setup := func() (queryCache *QueryCache) {
		queryCache = NewQueryCache(mockedTTL, f.FrozenTime)
		expectEmptyCache(queryCache) // Preliminary Check to weed out possible false positives.
		queryCache.entry["key1"] = createCacheEntry("value_1", f.FrozenTime.Add(-time.Hour))
		return queryCache
	}

	t.Run("Successful Creation of QueryCache", func(t *testing.T) {
		t.Run("should create a new non-nil instance successfully.", func(t *testing.T) {
			if queryCache := setup(); queryCache == nil {
				t.Error("Expected QueryCache to be defined, but got nil")
			}
		})

		t.Run("should create a new instance with a non-nil entry map,", func(t *testing.T) {
			if queryCache := setup(); queryCache.entry == nil {
				t.Error("Expected non-nil entry map, got nil")
			}
		})

		t.Run("should create a new instance with a non-nil mutex.", func(t *testing.T) {
			if queryCache := setup(); queryCache.mutex == nil {
				t.Error("Expected non-nil mutex, got nil")
			}
		})

		t.Run("should start the eviction loop successfully.", func(t *testing.T) {
			queryCache := setup()
			waitForEvictionLoop()
			if len(queryCache.entry) != 0 {
				t.Error("Expected cache to be empty, but found one or more entries")
			}
		})
	})
}
