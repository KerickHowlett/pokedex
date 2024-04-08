package query_cache

import "time"

// Save adds a new entry to the query cache with the specified key and value.
//
// The function acquires a lock on the query cache.
//
// Parameters:
//   - key: The key to associate with the value.
//   - value: The value to be stored in the cache.
//   - now: The current time to be used as the reference time for the cache. If not provided, the current system time will be used.
//
// Example:
//
//	 response, err := http.Get("https://example.com/api/v1")
//		cache.Save("key", response.Body)
func (qc *QueryCache) Save(key string, value []byte, now ...time.Time) {
	qc.mutex.Lock()
	defer qc.mutex.Unlock()

	createdAt := time.Now()
	if len(now) >= 1 {
		createdAt = now[0]
	}

	qc.entry[key] = cacheEntry{
		createdAt: &createdAt,
		value:     value,
	}
}
