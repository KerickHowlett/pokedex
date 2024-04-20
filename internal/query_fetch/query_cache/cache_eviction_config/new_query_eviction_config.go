package query_cache_config

import (
	"time"

	"query_fetch/query_cache/ttl"
)

// NewQueryEvictionConfig creates a new QueryEvictionConfig instance with the provided options.
//
// The default TTL is set to ttl.OneDay.
// The default Now function is set to use time.Now.
//
// Parameters:
//   - options: The options to set on the QueryEvictionConfig instance.
//
// Returns:
//   - *QueryEvictionConfig: The new QueryEvictionConfig instance.
//
// Example:
//
//	config := NewQueryEvictionConfig()
func NewQueryEvictionConfig(options ...QueryEvictionConfigOption) *QueryEvictionConfig {
	config := &QueryEvictionConfig{
		TTL: ttl.OneDay,
		Now: time.Now,
	}
	for _, option := range options {
		option(config)
	}
	return config
}
