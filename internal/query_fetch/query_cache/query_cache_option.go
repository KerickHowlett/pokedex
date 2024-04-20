package query_cache

import qec "query_fetch/query_cache/cache_eviction_config"

// QueryCacheOption is a function that sets an option on a QueryCache instance.
type QueryCacheOption func(*QueryCache)

// WithEvictionConfig sets the eviction configuration for the cache.
//
// Parameters:
//   - ec: The eviction configuration to set.
//
// Returns:
//   - A QueryCacheOption function that sets the eviction configuration for the cache.
//
// Example usage:
//
//	cache := NewQueryCache(WithEvictionConfig(&qec.QueryEvictionConfig{
//		TTL: 24 * time.Hour,
//	}))
func WithEvictionConfig(ec *qec.QueryEvictionConfig) QueryCacheOption {
	return func(qc *QueryCache) {
		qc.ec = ec
	}
}
