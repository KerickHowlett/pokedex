package query_cache

import "time"

// evictionLoop is responsible for evicting expired entries from the query cache.
//
// The function acquires a lock on the query cache, iterates over all entries,
// and deletes any entry that has expired based on the TTL and current time.
//
// Parameters:
//   - ttl: The Time-to-Live (TTL) duration that determines how long each entry will be kept in the cache before being evicted.
//   - now: The current time to be used as the reference time for the cache. If not provided, the current system time will be used.
//
// Example:
//
//	 ttl := 5 * time.Minute
//	 cache := NewQueryCache(ttl)
//		cache.evictionLoop(ttl)
func (qc *QueryCache) evictionLoop(ttl time.Duration, now ...time.Time) {
	qc.mutex.Lock()
	defer qc.mutex.Unlock()

	for key, entry := range qc.entry {
		if qc.isEntryExpired(&entry, ttl, now...) {
			delete(qc.entry, key)
		}
	}
}
