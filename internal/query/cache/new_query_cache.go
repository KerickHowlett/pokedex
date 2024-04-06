package query_cache

import "sync"

func NewQueryCache() *QueryCache {
	return &QueryCache{
		entry: make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}
}
