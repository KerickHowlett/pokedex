package querystate

// NewQueryState creates a new instance of QueryState with the specified options.
//
// Parameters:
//   - options: The options to configure the QueryState.
//
// Returns:
//   - A new instance of QueryState.
//
// Example:
//
//	state := NewQueryState(
//		WithNextURL("https://example.com/query?after=123"),
//		WithPreviousURL("https://example.com/query?before=123"),
//      WithResults("result1"),
//      WithResults("result2"),
// )
func NewQueryState[TResult any](options ...QueryStateOption[TResult]) (queryState *QueryState[TResult]) {
	queryState = &QueryState[TResult]{
		NextURL:     nil,
		PreviousURL: nil,
		Results:     []TResult{},
	}

	for _, option := range options {
		option(queryState)
	}

	return queryState
}
