package query_cache_config

import "time"

type QueryCacheConfigOption func(*QueryCacheConfig)

// WithNow sets the function that returns the current time,
// which is used for controlling the expiration time of the
// cache entries.
//
// Parameters:
//   - now: A function that returns the current time.
//
// Returns:
//   - A QueryCacheConfigOption function that sets the Now field of the QueryCacheConfig struct.
//
// Example usage:
//
//	cache := NewQueryCache(
//		WithTTL(5 * time.Minute),
//		WithNow(func() time.Time { return time.Now() }),
//	)
func WithNow(now func() time.Time) QueryCacheConfigOption {
	return func(qc *QueryCacheConfig) {
		qc.Now = now
	}
}

// WithTTL sets the Time-to-Live (TTL) duration for the QueryCache instance.
//
// Parameters:
//   - ttl: The Time-to-Live (TTL) duration for the QueryCache instance.
//
// Returns:
//   - A QueryCacheConfigOption function that sets the TTL field of the QueryCacheConfig struct.
//
// Example usage:
//
//	cache := NewQueryCache(
//		WithTTL(5 * time.Minute),
//	)
func WithTTL(ttl time.Duration) QueryCacheConfigOption {
	return func(qc *QueryCacheConfig) {
		qc.TTL = ttl
	}
}
