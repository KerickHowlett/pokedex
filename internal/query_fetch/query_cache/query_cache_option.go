package query_cache

import "time"

// QueryCacheOption is a function that sets an option on a QueryCache instance.
type QueryCacheOption func(*QueryCache)

// WithNow sets the current time to be used as the reference time for the cache.
//
// Parameters:
//   - now: The current time to be used as the reference time for the cache.
//
// Returns:
//   - A QueryCacheOption function that sets the current time to be used as the reference time for the cache.
//
// Example usage:
//
//	cache := NewQueryCache(WithNow(time.Now()))
func WithNow(now func() time.Time) QueryCacheOption {
	return func(qc *QueryCache) {
		qc.now = now
	}
}

// WithTTL sets the Time-to-Live (TTL) duration that determines how long each entry will be kept in the cache before being evicted.
//
// Parameters:
//   - ttl: The Time-to-Live (TTL) duration.
//
// Returns:
//   - A QueryCacheOption function that sets the Time-to-Live (TTL) duration.
//
// Example usage:
//
//	cache := NewQueryCache(WithTTL(ttl.OneDay))
func WithTTL(ttl time.Duration) QueryCacheOption {
	return func(qc *QueryCache) {
		qc.ttl = ttl
	}
}
