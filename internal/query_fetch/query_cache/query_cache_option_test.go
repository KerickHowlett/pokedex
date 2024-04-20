package query_cache

import (
	qec "query_fetch/query_cache/cache_eviction_config"
	"testing"
)

func TestWithEvictionConfig(t *testing.T) {
	setup := func() (actual *qec.QueryEvictionConfig, expected *qec.QueryEvictionConfig) {
		expected = &qec.QueryEvictionConfig{}
		queryCache := &QueryCache{}
		WithEvictionConfig(expected)(queryCache)
		actual = queryCache.ec
		return actual, expected
	}

	t.Run("should set the eviction config field.", func(t *testing.T) {
		if actual, expected := setup(); actual != expected {
			t.Errorf("Expected: %v\nActual: %v", expected, actual)
		}
	})
}
