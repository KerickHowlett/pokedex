package querystate

type result struct {
	Name string
}

type queryTestState = QueryState[result]
