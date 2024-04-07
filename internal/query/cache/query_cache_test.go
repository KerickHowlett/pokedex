package query_cache

import "time"

const (
	mockedTTL   = time.Microsecond
	invalid_key = "invalid_key"
	key         = "cache_index_key"
	value       = "cached_value"
)

func createCacheEntry(value string, createdAt time.Time) cacheEntry {
	return cacheEntry{
		createdAt: &createdAt,
		value:     []byte(value),
	}
}
