package query_cache

import "testing"

func TestNewQueryCache(t *testing.T) {
	setup := func() *QueryCache {
		return NewQueryCache(mockedInterval)
	}

	t.Run("Successful Creation of QueryCache", func(t *testing.T) {
		t.Run("should create a new non-nil instance successfully", func(t *testing.T) {
			if queryCache := setup(); queryCache == nil {
				t.Error("Expected QueryCache to be defined, but got nil")
			}
		})

		t.Run("should create a new instance with a non-nil entry map", func(t *testing.T) {
			if queryCache := setup(); queryCache.entry == nil {
				t.Error("Expected non-nil entry map, got nil")
			}
		})

		t.Run("should create a new instance with a non-nil mutex", func(t *testing.T) {
			if queryCache := setup(); queryCache.mutex == nil {
				t.Error("Expected non-nil mutex, got nil")
			}
		})
	})
}
