package query_cache

import (
	"testing"
	"time"

	f "test_tools/fixtures"
)

func TestIsEntryExpired(t *testing.T) {
	setup := func(createdAt time.Time) (queryCache *QueryCache, entry *cacheEntry, ttl time.Duration) {
		queryCache = &QueryCache{}
		entry = &cacheEntry{
			createdAt: &createdAt,
		}
		ttl = time.Hour

		return queryCache, entry, ttl
	}

	t.Run("should return TRUE if entry has expired.", func(t *testing.T) {
		cache, entry, ttl := setup(f.FrozenTime)
		if expired := cache.isEntryExpired(entry, ttl, f.FrozenTime); !expired {
			t.Errorf("Expected entry to be expired, but it is not")
		}
	})

	t.Run("should return FALSE if entry has NOT expired.", func(t *testing.T) {
		cache, entry, ttl := setup(f.FrozenTime.Add(time.Hour * 24))
		if expired := cache.isEntryExpired(entry, ttl, f.FrozenTime); expired {
			t.Errorf("Expected entry to NOT be expired, but it is")
		}
	})
}
