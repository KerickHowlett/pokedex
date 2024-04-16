package query_cache

import (
	"testing"
	"time"

	"query_fetch/query_cache/ttl"
	f "test_tools/fixtures"
)

func TestWithNow(t *testing.T) {
	runWithNowTests := func() *QueryCache {
		cache := &QueryCache{}
		WithNow(func() time.Time { return f.FrozenTime })(cache)
		return cache
	}

	t.Run("should set the QueryCache.now field with the provided time.", func(t *testing.T) {
		if cache := runWithNowTests(); cache.now() != f.FrozenTime {
			t.Errorf("Expected QueryCache.now to be %s, but got %s", f.FrozenTime.String(), cache.now().String())
		}
	})
}

func TestWithTTL(t *testing.T) {
	runWithTTLTests := func() *QueryCache {
		cache := &QueryCache{}
		WithTTL(ttl.Disable)(cache)
		return cache
	}

	t.Run("should set the cache time-to-live (TTL) with the 'disabled' setting provided.", func(t *testing.T) {
		if cache := runWithTTLTests(); cache.ttl != ttl.Disable {
			t.Errorf("Expected ttl to be %v, but got %v", ttl.OneDay, cache.ttl)
		}
	})
}
