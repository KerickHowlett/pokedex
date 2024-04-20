package query_cache

import (
	qec "query_fetch/query_cache/cache_eviction_config"
	"sync"
)

// NewQueryCache creates a new instance of QueryCache with the specified Time-to-Live (TTL) duration.
//
// Parameters:
//   - options: A variadic list of QueryCacheOption functions to configure the QueryCache instance.
//
// Returns:
//   - A new instance of QueryCache.
//
// Example:
//
//	cache := NewQueryCache(5 * time.Minute)
func NewQueryCache(options ...QueryCacheOption) *QueryCache {
	qc := &QueryCache{
		entry: make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
		ec:    qec.NewQueryEvictionConfig(),
	}

	for _, option := range options {
		option(qc)
	}

	go qc.evictionLoop()

	return qc
}
