package query_cache

import "time"

func (qc *QueryCache) isEntryExpired(entry *cacheEntry, ttl time.Duration, now ...time.Time) bool {
	currentTime := time.Now()
	if len(now) > 0 {
		currentTime = now[0]
	}

	return entry.createdAt.Add(-ttl).Before(currentTime)
}
