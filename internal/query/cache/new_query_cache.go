package query_cache

import (
	"sync"
	"time"
)

func NewQueryCache(interval time.Duration, now ...time.Time) *QueryCache {
	qc := &QueryCache{
		entry: make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}

	go qc.evictionLoop(interval, now...)

	return qc
}
