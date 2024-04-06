package querystate

// QueryState represents the state of a query result.
//
// Generic Constraints:
//
//	T - The query result's type, whether it'd be a primitive type or a struct.
//
// Fields:
//   - NextURL: Represents the URL for the next page of results. Nullable.
//   - PreviousURL: Represents the URL of the previous page of results. Nullable.
//   - Results: A slice of type T that represents the results of a query state.
//     It is populated with data from the JSON field "results".
//
// Example:
//
//	state := &QueryState{}
type QueryState[T any] struct {
	NextURL     *string `json:"next"`
	PreviousURL *string `json:"previous"`
	Results     []T     `json:"results"`
}
