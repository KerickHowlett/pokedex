package query_fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	qc "query_fetch/query_cache"
	qec "query_fetch/query_cache/cache_eviction_config"
)

// QueryFetchFunc is used to type function parameters that requires a function
// matching QueryFetch's call signature in order to execute properly.
//
// Generic Constraints:
//   - TQuery: The struct of the query payload.
//
// Parameters:
//   - url: The URL to fetch the query from.
//   - ttlOption: An optional time.Duration parameter that sets the time-to-live for the query cache. The default value is 24 Hours. If TTL is set to zero (0), the cache will not be used.
//
// Returns:
//   - An error if the query fetch operation fails.
type QueryFetchFunc[TQuery any] func(url string, config ...*qec.QueryEvictionConfig) (query *TQuery, err error)

// QueryFetch sends an HTTP GET request to the specified URL and fetches the
// query result into the provided query state.
//
// Generic Constraints:
//   - TQuery: The type of the query result, whether it'd be a primitive type or a struct.
//
// Parameters:
//   - url: The URL to fetch the query from.
//   - query: A pointer to the query struct to be populated with the fetched HTTP Response.
//   - config: An optional QueryEvictionConfig parameter that sets the eviction configuration for the query cache.
//
// Returns:
//   - A pointer to the query struct populated with the fetched HTTP Response.
//   - An error if the query fetch operation fails.
//
// Example:
//
//	query, err := QueryFetch("https://example.com/query", query)
func QueryFetch[TQuery any](url string, config ...*qec.QueryEvictionConfig) (query *TQuery, err error) {
	evictionConfig := qec.NewQueryEvictionConfig()
	if len(config) > 0 {
		evictionConfig = config[0]
	}

	cache := qc.NewQueryCache(qc.WithEvictionConfig(evictionConfig))
	if cachedResponse, cacheHit := cache.Find(url); cacheHit {
		return decode[TQuery](cachedResponse)
	}

	response, err := http.Get(url)
	if err != nil {
		return query, fmt.Errorf("error while attempting to fetch query: %v", err)
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return query, fmt.Errorf("error from reading response body: %v", err)
	}

	if !isSuccessfulResponse(response.StatusCode) {
		return query, fmt.Errorf("error with response: %s", body)
	}

	if _, err = decode[TQuery](body); err != nil {
		return query, err
	}

	cache.Save(url, body)

	return query, nil
}

// decode parses the HTTP Response's JSON body and populates the
// payload it's values.
//
// Generic Constraints:
//   - TQuery: The schematic of the intended query payload.
//
// Parameters:
//   - body: The HTTP Response's body.
//
// Returns:
//   - A pointer to the query struct populated with the fetched HTTP Response.
//   - An error if the response body cannot be parsed.
//
// Example usage:
//
//	decodedQuery, err := decode(body, query)
func decode[TQuery any](body []byte) (query *TQuery, err error) {
	if err := json.Unmarshal(body, &query); err != nil {
		return query, fmt.Errorf("error with parsing payload %v", err)
	}
	return query, nil
}

// isSuccessfulResponse checks if the given HTTP status code indicates a successful response.
//
// It returns true if the status code is in the range of 200 to 299 (inclusive), and false otherwise.
//
// Parameters:
//
//	statusCode: The HTTP status code to check.
//
// Returns:
//   - bool: A boolean value indicating if the status code is successful.
//
// Example usage:
//
//	isSuccessful := isSuccessfulResponse(http.StatusOK)
func isSuccessfulResponse(statusCode int) bool {
	return statusCode >= http.StatusOK &&
		statusCode < http.StatusMultipleChoices
}
