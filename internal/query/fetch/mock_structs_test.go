package query_fetch

type testResult struct {
	Name string
}

type testQuery struct {
	NextURL     string
	PreviousURL string
	Results     []testResult
}
