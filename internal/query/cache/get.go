package query_cache

// Get retrieves the value associated with the given key from the query cache.
//
// Parameters:
//   - key: The key to retrieve the value for.
//
// Returns:
//   - The value associated with the key, or nil if the key is not found.
//   - A boolean value indicating whether the key was found in the cache.
//
// Example:
//
//	value, cacheHit := cache.Get("key")
//	if !cacheHit {
//	    log.Println("key not found in cache")
//	}
func (qc *QueryCache) Get(key string) (value []byte, entryFound bool) {
	qc.mutex.Lock()
	defer qc.mutex.Unlock()

	entry, entryFound := qc.entry[key]
	if !entryFound {
		return nil, false
	}

	return entry.value, true
}
