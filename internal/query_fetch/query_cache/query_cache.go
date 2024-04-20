package query_cache

import (
	"sync"
	"time"

	qec "query_fetch/query_cache/cache_eviction_config"
)

// cacheEntry represents a single entry in the cache.
//
// Fields:
//   - createdAt: The time when the cache entry was created.
//   - value: The cached value.
type cacheEntry struct {
	// createdAt is the time when the cache entry was created.
	createdAt time.Time
	// value is the cached value.
	value []byte
}

// QueryCache is a cache that stores HTTP query results.
//
// The cache is thread-safe and can be accessed concurrently.
//
// Fields:
//   - entry: A map that stores one or more cached entries.
//   - mutex: A mutex that synchronizes access to the cache and makes the struct thread-safe.
//   - now: The current time to be used as the reference time for the cache.
type QueryCache struct {
	// entry is a map that stores one or more cached entries.
	entry map[string]cacheEntry
	// mutex synchronizes access to the cache and makes the struct thread-safe.
	mutex *sync.Mutex
	// Eviction Config settings for the cache.
	ec *qec.QueryEvictionConfig
}

// Find retrieves the value associated with the given key from the query cache.
//
// Parameters:
//   - key: The key to retrieve the value for.
//
// Returns:
//   - The value associated with the key, or nil if the key is not found.
//   - A boolean value indicating whether the key was found in the cache.
//
// Example:
//
//	value, cacheHit := cache.Find("key")
//	if !cacheHit {
//	    log.Println("key not found in cache")
//	}
func (qc *QueryCache) Find(key string) (value []byte, entryFound bool) {
	qc.mutex.Lock()
	defer qc.mutex.Unlock()

	if entry, entryFound := qc.entry[key]; entryFound {
		return entry.value, true
	}

	return nil, false
}

// Save adds a new entry to the query cache with the specified key and value.
//
// The function acquires a lock on the query cache.
//
// Parameters:
//   - key: The key to associate with the value.
//   - value: The value to be stored in the cache.
//
// Example:
//
// response, err := http.Get("https://example.com/api/v1")
// cache.Save("key", response.Body)
func (qc *QueryCache) Save(key string, value []byte) {
	qc.mutex.Lock()
	defer qc.mutex.Unlock()

	qc.entry[key] = cacheEntry{
		createdAt: qc.ec.Now(),
		value:     value,
	}
}

// evictionLoop is responsible for evicting expired entries from the query cache.
//
// The function acquires a lock on the query cache, iterates over all entries,
// and deletes any entry that has expired based on the TTL and current time.
//
// Example usage:
//
// cache := NewQueryCache()
// go cache.evictionLoop()
func (qc *QueryCache) evictionLoop() {
	qc.mutex.Lock()
	defer qc.mutex.Unlock()

	for key, entry := range qc.entry {
		if qc.isEntryExpired(&entry) {
			delete(qc.entry, key)
		}
	}
}

// isEntryExpired checks if the cache entry is expired based on the Time-to-Live (TTL) duration.
//
// Parameters:
//   - entry: The cache entry to check for expiration.
//
// Returns:
//   - A boolean value indicating whether the cache entry is expired.
//
// Example usage:
//
//	isExpired := cache.isEntryExpired(entry, 5 * time.Minute)
func (qc *QueryCache) isEntryExpired(entry *cacheEntry) bool {
	return entry.createdAt.Add(-qc.ec.TTL).Before(qc.ec.Now())
}
