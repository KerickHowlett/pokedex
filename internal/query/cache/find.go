package query_cache

func (qc *QueryCache) Find(key string) (value []byte, exists bool) {
	qc.mutex.Lock()
	defer qc.mutex.Unlock()

	if entry, exists := qc.entry[key]; exists {
		return entry.value, true
	}

	return nil, false
}
