package query_cache

import (
	"testing"
)

func TestQueryCache_Find(t *testing.T) {
	qc := NewQueryCache()

	// Add a sample entry to the cache
	key := "testKey"
	value := []byte("testValue")
	qc.entry[key] = cacheEntry{value: value}

	t.Run("when cache entry exists in QueryCache", func(t *testing.T) {
		t.Run("should return with a boolean value of TRUE", func(t *testing.T) {
			if _, exists := qc.Find(key); !exists {
				t.Errorf("Expected key '%s' to exist in the cache", key)
			}
		})

		t.Run("should return the value of the cache entry", func(t *testing.T) {
			if result, _ := qc.Find(key); string(result) != string(value) {
				t.Errorf("Expected value '%s' for key '%s', but got '%s'", string(value), key, string(result))
			}
		})
	})

	t.Run("when cache entry does NOT exist in QueryCache", func(t *testing.T) {
		t.Run("should return with a boolean value of FALSE", func(t *testing.T) {
			if _, exists := qc.Find("nonExistingKey"); exists {
				t.Errorf("Expected key 'nonExistingKey' to NOT exist in the cache")
			}
		})

		t.Run("should return nil as the value", func(t *testing.T) {
			if result, _ := qc.Find("nonExistingKey"); result != nil {
				t.Errorf("Expected nil value for non-existing key, but got '%s'", string(result))
			}
		})
	})
}
