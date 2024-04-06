package querystate

// result represents the result of a query state operation.
//
// Fields:
//   - Name: The name of the result.
type result struct {
	Name string
}

// queryTestState is a type alias for QueryState[result].
// It represents the state of a query test and is used in
// testing scenarios.
type queryTestState = QueryState[result]
