package query_cache

import "testing"

func TestNewQueryCache(t *testing.T) {
	t.Run("Successful Creation of QueryCache", func(t *testing.T) {
		t.Run("should create a new non-nil instance successfully", func(t *testing.T) {
			if queryCache := NewQueryCache(); queryCache == nil {
				t.Error("Expected QueryCache to be defined, but got nil")
			}
		})

		t.Run("should create a new instance with a non-nil entry map", func(t *testing.T) {
			if queryCache := NewQueryCache(); queryCache.entry == nil {
				t.Error("Expected non-nil entry map, got nil")
			}
		})

		t.Run("should create a new instance with a non-nil mutex", func(t *testing.T) {
			if queryCache := NewQueryCache(); queryCache.mutex == nil {
				t.Error("Expected non-nil mutex, got nil")
			}
		})
	})
}
