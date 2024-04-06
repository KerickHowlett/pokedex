package querystate

// QueryStateOption is a functional option type for configuring a QueryState.
// It is used to modify the behavior of a QueryState instance.
//
// Generic Constraints:
//
//   - TResult: The type of the query result, whether it'd be a primitive type or a struct.
type QueryStateOption[TResult any] func(*QueryState[TResult])
