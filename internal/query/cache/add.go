package query_cache

import "time"

func (qc *QueryCache) Add(key string, value []byte, now ...time.Time) {
	qc.mutex.Lock()
	defer qc.mutex.Unlock()

	createdAt := time.Now()
	if len(now) >= 1 {
		createdAt = now[0]
	}

	qc.entry[key] = cacheEntry{
		createdAt: &createdAt,
		value:     value,
	}
}
