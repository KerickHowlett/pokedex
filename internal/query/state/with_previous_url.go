package querystate

// WithPreviousURL sets the next URL for the query state.
//
// It is a QueryStateOption function that takes a nullable URL pointer and returns
// a function that sets the PreviousURL field of the QueryState.
//
// Generic Constraint:
//
// - TResult: The type of the query result, whether it'd be a primitive type or a struct.
//
// Parameters:
//   - url: The URL to set as the next URL.
//
// Returns:
//   - A function that sets the PreviousURL field of the QueryState.
//
// Example:
//
//	state := NewQueryState(
//		WithPreviousURL("https://example.com/query?after=123"),
//	)
func WithPreviousURL[TResult any](url *string) QueryStateOption[TResult] {
	return func(m *QueryState[TResult]) {
		m.PreviousURL = url
	}
}
