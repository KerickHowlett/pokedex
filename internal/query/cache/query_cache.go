package query_cache

import "sync"

type QueryCache struct {
	entry map[string]cacheEntry
	mutex *sync.Mutex
}
