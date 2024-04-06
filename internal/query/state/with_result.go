package querystate

// WithResult is a QueryStateOption function that appends the given result to
// the query state's Results slice.
//
// It takes a result of type TResult and returns a function that modifies the
// QueryState[TResult] instance.
//
// Generic Constraints:
//
// - TResult: The type of the query result, whether it'd be a primitive type or a struct.
//
// Parameters:
//   - result: The result to append to the query state's Results slice.
//
// Returns:
//   - A function that appends the given result to the query state's Results slice.
//
// Example:
//
//	state := NewQueryState(
//		WithResult("result1"),
//		WithResult("result2"),
//	)
func WithResult[TResult any](result TResult) QueryStateOption[TResult] {
	return func(m *QueryState[TResult]) {
		m.Results = append(m.Results, result)
	}
}
