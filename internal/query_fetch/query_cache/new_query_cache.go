package query_cache

import (
	"sync"
	"time"
)

// NewQueryCache creates a new instance of QueryCache with the specified Time-to-Live (TTL) duration.
//
// Parameters:
//   - ttl: The Time-to-Live (TTL) that determines how long each entry will be kept in the cache before being evicted.
//   - now: The current time to be used as the reference time for the cache. If not provided, the current system time will be used.
//
// Returns:
//   - A new instance of QueryCache.
//
// Example:
//
//	cache := NewQueryCache(5 * time.Minute)
func NewQueryCache(ttl time.Duration, now ...time.Time) *QueryCache {
	qc := &QueryCache{
		entry: make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}

	go qc.evictionLoop(ttl, now...)

	return qc
}
