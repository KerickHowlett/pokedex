package query_cache

import "sync"

// QueryCache is a cache that stores HTTP query results.
//
// The cache is thread-safe and can be accessed concurrently.
//
// Fields:
//   - entry: A map that stores one or more cached entries.
//   - mutex: A mutex that synchronizes access to the cache and makes the struct thread-safe.
type QueryCache struct {
	entry map[string]cacheEntry
	mutex *sync.Mutex
}
