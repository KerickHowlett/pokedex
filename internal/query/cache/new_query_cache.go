package query_cache

import (
	"sync"
	"time"
)

func NewQueryCache(ttl time.Duration, now ...time.Time) *QueryCache {
	qc := &QueryCache{
		entry: make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}

	go qc.evictionLoop(ttl, now...)

	return qc
}
