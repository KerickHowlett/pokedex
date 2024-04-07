package query_cache

func (qc *QueryCache) Get(key string) (value []byte, entryFound bool) {
	qc.mutex.Lock()
	defer qc.mutex.Unlock()

	entry, entryFound := qc.entry[key]
	if !entryFound {
		return nil, false
	}

	return entry.value, true
}
