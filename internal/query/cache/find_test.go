package query_cache

import "testing"

func TestQueryCache_Find(t *testing.T) {
	setup := func() (queryCache *QueryCache) {
		queryCache = NewQueryCache(mockedTTL)
		queryCache.entry[key] = cacheEntry{value: []byte(value)}

		return queryCache
	}

	t.Run("when cache entry exists in QueryCache", func(t *testing.T) {
		t.Run("should return with a boolean value of TRUE", func(t *testing.T) {
			queryCache := setup()
			if _, exists := queryCache.Find(key); !exists {
				t.Errorf("Expected key '%s' to exist in the cache", key)
			}
		})

		t.Run("should return the value of the cache entry", func(t *testing.T) {
			queryCache := setup()
			if result, _ := queryCache.Find(key); string(result) != string(value) {
				t.Errorf("Expected value '%s' for key '%s', but got '%s'", string(value), key, string(result))
			}
		})
	})

	t.Run("when cache entry does NOT exist in QueryCache", func(t *testing.T) {
		t.Run("should return with a boolean value of FALSE", func(t *testing.T) {
			queryCache := setup()
			if _, exists := queryCache.Find(invalid_key); exists {
				t.Errorf("Expected key '%s' to NOT exist in the cache", invalid_key)
			}
		})

		t.Run("should return nil as the value", func(t *testing.T) {
			queryCache := setup()
			if result, _ := queryCache.Find(invalid_key); result != nil {
				t.Errorf("Expected nil value for non-existing key, but got '%s'", string(result))
			}
		})
	})
}
