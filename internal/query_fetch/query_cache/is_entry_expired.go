package query_cache

import "time"

// isEntryExpired checks if the cache entry is expired based on the Time-to-Live (TTL) duration.
//
// Parameters:
//   - entry: The cache entry to check for expiration.
//   - ttl: The Time-to-Live (TTL) duration that determines how long each entry will be kept in the cache before being evicted.
//   - now: The current time to be used as the reference time for the cache. If not provided, the current system time will be used.
//
// Returns:
//   - A boolean value indicating whether the cache entry is expired.
//
// Example:
//
//	isExpired := cache.isEntryExpired(entry, 5 * time.Minute)
func (qc *QueryCache) isEntryExpired(entry *cacheEntry, ttl time.Duration, now ...time.Time) bool {
	currentTime := time.Now()
	if len(now) > 0 {
		currentTime = now[0]
	}

	return entry.createdAt.Add(-ttl).Before(currentTime)
}
