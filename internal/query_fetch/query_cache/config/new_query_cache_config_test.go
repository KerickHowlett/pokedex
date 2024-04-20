package query_cache_config

import (
	"testing"

	"query_fetch/query_cache/ttl"
)

func TestWithNewQueryCacheConfig(t *testing.T) {
	setup := func() (actual *QueryCacheConfig) {
		actual = NewQueryCacheConfig()
		return actual
	}

	t.Run("should return a new QueryCacheConfig instance.", func(t *testing.T) {
		if actual := setup(); actual == nil {
			t.Error("Expected a new QueryCacheConfig instance, but it was nil.")
		}
	})

	t.Run("should set the default TTL field of the QueryCacheConfig struct.", func(t *testing.T) {
		if actual := setup(); actual.TTL != ttl.OneDay {
			t.Errorf("Expected: %v,\nActual: %v", ttl.OneDay, actual.TTL)
		}
	})

	t.Run("should set the default Now field of the QueryCacheConfig struct.", func(t *testing.T) {
		if actual := setup(); actual.Now == nil {
			t.Error("Expected the Now field to be set, but it was nil.")
		}
	})
}
