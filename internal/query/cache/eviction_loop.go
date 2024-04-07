package query_cache

import "time"

func (qc *QueryCache) evictionLoop(ttl time.Duration, now ...time.Time) {
	qc.mutex.Lock()
	defer qc.mutex.Unlock()

	for key, entry := range qc.entry {
		if qc.isEntryExpired(&entry, ttl, now...) {
			delete(qc.entry, key)
		}
	}
}
