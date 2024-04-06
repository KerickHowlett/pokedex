package query_cache

import "time"

type cacheEntry struct {
	createdAt *time.Time
	value     []byte
}
