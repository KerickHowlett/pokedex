package query_eviction_config

import "time"

type QueryEvictionConfigOption func(*QueryEvictionConfig)

// WithNow sets the function that returns the current time,
// which is used for controlling the expiration time of the
// cache entries.
//
// Parameters:
//   - now: A function that returns the current time.
//
// Returns:
//   - A QueryEvictionConfigOption function that sets the Now field of the QueryEvictionConfig struct.
//
// Example usage:
//
//	cache := NewQueryCache(
//		WithTTL(5 * time.Minute),
//		WithNow(func() time.Time { return time.Now() }),
//	)
func WithNow(now func() time.Time) QueryEvictionConfigOption {
	return func(qc *QueryEvictionConfig) {
		qc.Now = now
	}
}

// WithTTL sets the Time-to-Live (TTL) duration for the QueryCache instance.
//
// Parameters:
//   - ttl: The Time-to-Live (TTL) duration for the QueryCache instance.
//
// Returns:
//   - A QueryEvictionConfigOption function that sets the TTL field of the QueryEvictionConfig struct.
//
// Example usage:
//
//	cache := NewQueryCache(
//		WithTTL(5 * time.Minute),
//	)
func WithTTL(ttl time.Duration) QueryEvictionConfigOption {
	return func(qc *QueryEvictionConfig) {
		qc.TTL = ttl
	}
}
