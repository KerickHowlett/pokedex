package query_fetch

import (
	"fmt"
	"io"
	"net/http"
	"time"

	qc "query/cache"
	qs "query/state"
)

// QueryFetch sends an HTTP GET request to the specified URL and fetches the
// query result into the provided query state.
//
// Generic Constraints:
//
//   - TResult: The type of the query result, whether it'd be a primitive type or a struct.
//
// Parameters:
//   - url: The URL to fetch the query from.
//   - queryState: The query state to fetch the query result into.
//
// Returns:
//   - An error if the query fetch operation fails.
//
// Example:
//
//	state := &QueryState{}
//	err := QueryFetch("https://example.com/query", state)
//	if err != nil {
//	    log.Fatalf("error while fetching query: %v", err)
//	}
func QueryFetch[TResult any](url string, queryState *qs.QueryState[TResult], ttlOption ...time.Duration) error {
	ttl := oneDayTTL
	if len(ttlOption) > 0 {
		ttl = ttlOption[0]
	}

	cache := qc.NewQueryCache(ttl)
	if cachedResponse, cacheHit := cache.Find(url); cacheHit {
		return parseQueryFetchResponse(cachedResponse, queryState)
	}

	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while attempting to fetch query: %v", err)
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return fmt.Errorf("error from reading response body: %v", err)
	}

	if !isSuccessfulResponse(response.StatusCode) {
		return fmt.Errorf("error with response: %s", body)
	}

	if err = parseQueryFetchResponse(body, queryState); err != nil {
		return err
	}

	cache.Save(url, body)

	return nil
}
