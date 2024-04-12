package query_fetch

import (
	"fmt"
	"io"
	"net/http"
	"time"

	qc "query/cache"
	qttl "query/fetch/ttl"
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
//   - queryState: A pointer to the query state to populate with the fetched query result.
//   - ttlOption: An optional time.Duration parameter that sets the time-to-live for the query cache. The default value is 24 Hours. If TTL is set to zero (0), the cache will not be used.
//
// Returns:
//   - An error if the query fetch operation fails.
//
// Example:
//
//	err := QueryFetch("https://example.com/query", payload)
//	if err != nil {
//	    log.Fatalf("error while fetching query: %v", err)
//	}
func QueryFetch[TPayload any](url string, payload *TPayload, ttlOption ...time.Duration) error {
	ttl := qttl.OneDayTTL
	if len(ttlOption) > 0 {
		ttl = ttlOption[0]
	}

	cache := qc.NewQueryCache(ttl)
	if cachedResponse, cacheHit := cache.Find(url); cacheHit {
		return decode(cachedResponse, payload)
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

	if err = decode(body, payload); err != nil {
		return err
	}

	cache.Save(url, body)

	return nil
}
