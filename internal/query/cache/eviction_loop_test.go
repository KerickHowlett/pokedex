package query_cache

import (
	"strconv"
	"strings"
	"testing"
	"time"

	f "test_tools/fixtures"
)

func TestQueryCache_evictionLoop(t *testing.T) {
	setup := func() (queryCache *QueryCache) {
		queryCache = NewQueryCache(mockedTTL)

		queryCache.entry["key1"] = createCacheEntry("value_1", f.FrozenTime.Add(-time.Hour))
		queryCache.entry["key2"] = createCacheEntry("value_2", f.FrozenTime.Add(-time.Minute))
		queryCache.entry["key3"] = createCacheEntry("value_3", f.FrozenTime.Add(time.Hour))

		return queryCache
	}

	expectCacheEntryExistence := func(key string, exists bool, cache *QueryCache) {
		existsLabel := strings.ToUpper(strconv.FormatBool(exists))
		if _, found := cache.entry[key]; found != exists {
			t.Errorf("Expected key '%s' to exist in the cache: %s, but it is not", key, existsLabel)
		}
	}

	expectEntryCount := func(length int, cache *QueryCache) {
		if entryCount := len(cache.entry); entryCount != length {
			t.Errorf("Expected cache to have %d entries, but only found %d instead", length, entryCount)
		}
	}

	t.Run("should only remove the expired cache entries.", func(t *testing.T) {
		cache := setup()

		// Preliminary Check to weed out possible false positives.
		expectEntryCount(3, cache)
		cache.evictionLoop(time.Minute, f.FrozenTime)

		expectEntryCount(1, cache)
		expectCacheEntryExistence("key1", false, cache)
		expectCacheEntryExistence("key2", false, cache)
		expectCacheEntryExistence("key3", true, cache)
	})
}
