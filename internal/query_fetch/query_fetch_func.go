package query_fetch

import "time"

// QueryFetchFunc is used to type function parameters that requires a function
// matching QueryFetch's call signature in order to execute properly.
//
// Generic Constraints:
//
//   - TQuery: The struct of the query payload.
type QueryFetchFunc[TQuery any] func(url string, query *TQuery, ttlOption ...time.Duration) error
