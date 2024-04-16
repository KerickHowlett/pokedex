package query_cache

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
	"time"

	f "test_tools/fixtures"
)

const (
	mockedTTL   = time.Microsecond
	invalid_key = "invalid_key"
	key         = "cache_index_key"
	value       = "cached_value"
)

func TestQueryCache_Find(t *testing.T) {
	setup := func() *QueryCache {
		queryCache := NewQueryCache(WithNow(func() time.Time { return f.FrozenTime }))
		queryCache.entry[key] = cacheEntry{value: []byte(value)}

		return queryCache
	}

	t.Run("Cache Hit", func(t *testing.T) {
		t.Parallel()
		t.Run("should return the value of the cache entry if it exists.", func(t *testing.T) {
			t.Parallel()
			cache := setup()
			if cachedValue, _ := cache.Find(key); !bytes.Equal(cachedValue, []byte(value)) {
				t.Errorf("Expected value to be %v, but got %v", value, cachedValue)
			}
		})

		t.Run("should return TRUE if the cache entry exists.", func(t *testing.T) {
			t.Parallel()
			cache := setup()
			if _, entryFound := cache.Find(key); !entryFound {
				t.Error("Expected cache entry to exist, but it does not.")
			}
		})
	})

	t.Run("Cache Miss", func(t *testing.T) {
		t.Parallel()
		t.Run("should return nil if the cache entry does not exist.", func(t *testing.T) {
			cache := setup()
			if cachedValue, _ := cache.Find(invalid_key); cachedValue != nil {
				t.Errorf("Expected value to be nil, but got %v", cachedValue)
			}
		})

		t.Run("should return FALSE if the cache entry does not exist.", func(t *testing.T) {
			t.Parallel()
			cache := setup()
			if _, entryFound := cache.Find(invalid_key); entryFound {
				t.Error("Expected cache entry to not exist, but it does.")
			}
		})
	})
}

func TestQueryCache_Save(t *testing.T) {
	runQueryCacheSaveTest := func() (cache *QueryCache) {
		now := func() time.Time { return f.FrozenTime }
		cache = NewQueryCache(
			WithNow(now),
			WithTTL(time.Hour),
		)
		cache.Save(key, []byte(value))
		return cache
	}

	t.Run("should save new cache entry.", func(t *testing.T) {
		t.Parallel()
		cache := runQueryCacheSaveTest()
		if _, exists := cache.entry[key]; !exists {
			t.Errorf("Expected entry with key '%s' to exist in cache, but it doesn't", key)
		}
	})

	t.Run("should set createdAt with the current time.", func(t *testing.T) {
		t.Parallel()
		cache := runQueryCacheSaveTest()
		if entry := cache.entry[key]; !entry.createdAt.Equal(f.FrozenTime) {
			t.Errorf("Expected createdAt to be %v, but got %v", f.FrozenTime.String(), entry.createdAt.String())
		}
	})

	t.Run("should set value with the provided value with it unmutated.", func(t *testing.T) {
		t.Parallel()
		cache := runQueryCacheSaveTest()
		if entry := cache.entry[key]; string(entry.value) != value {
			t.Errorf("Expected value to be '%s', but got '%s'", value, string(entry.value))
		}
	})
}

func TestQueryCache_evictionLoop(t *testing.T) {
	createCacheEntry := func(value string, createdAt time.Time) cacheEntry {
		return cacheEntry{
			createdAt: createdAt,
			value:     []byte(value),
		}
	}

	setup := func() (queryCache *QueryCache) {
		queryCache = NewQueryCache(
			WithNow(func() time.Time { return f.FrozenTime }),
			WithTTL(time.Microsecond),
		)

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
		cache.evictionLoop()

		expectEntryCount(1, cache)
		expectCacheEntryExistence("key1", false, cache)
		expectCacheEntryExistence("key2", false, cache)
		expectCacheEntryExistence("key3", true, cache)
	})
}

func TestIsEntryExpired(t *testing.T) {
	setup := func(createdAt time.Time) (queryCache *QueryCache, entry *cacheEntry) {
		queryCache = NewQueryCache(
			WithNow(func() time.Time { return f.FrozenTime }),
			WithTTL(time.Hour),
		)
		entry = &cacheEntry{createdAt: createdAt}

		return queryCache, entry
	}

	t.Run("should return TRUE if entry has expired.", func(t *testing.T) {
		t.Parallel()
		cache, entry := setup(f.FrozenTime)
		if expired := cache.isEntryExpired(entry); !expired {
			t.Errorf("Expected entry to be expired, but it is not")
		}
	})

	t.Run("should return FALSE if entry has NOT expired.", func(t *testing.T) {
		t.Parallel()
		cache, entry := setup(f.FrozenTime.Add(time.Hour * 24))
		if expired := cache.isEntryExpired(entry); expired {
			t.Errorf("Expected entry to NOT be expired, but it is")
		}
	})
}
