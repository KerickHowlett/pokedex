package query_cache

import "time"

// cacheEntry represents a single entry in the cache.
//
// Fields:
//   - createdAt: The time when the cache entry was created.
//   - value: The cached value.
type cacheEntry struct {
	createdAt *time.Time
	value     []byte
}
