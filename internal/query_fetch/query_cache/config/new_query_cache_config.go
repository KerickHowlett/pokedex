package query_cache_config

import (
	"time"

	"query_fetch/query_cache/ttl"
)

// NewQueryCacheConfig creates a new QueryCacheConfig instance with the provided options.
//
// The default TTL is set to ttl.OneDay.
// The default Now function is set to use time.Now.
//
// Parameters:
//   - options: The options to set on the QueryCacheConfig instance.
//
// Returns:
//   - *QueryCacheConfig: The new QueryCacheConfig instance.
//
// Example:
//
//	config := NewQueryCacheConfig()
func NewQueryCacheConfig(options ...QueryCacheConfigOption) *QueryCacheConfig {
	config := &QueryCacheConfig{
		TTL: ttl.OneDay,
		Now: time.Now,
	}
	for _, option := range options {
		option(config)
	}
	return config
}
