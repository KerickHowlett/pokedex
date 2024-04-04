package querystate

type QueryState[T any] struct {
	NextURL     *string `json:"next"`
	PreviousURL *string `json:"previous"`
	Results     []T     `json:"results"`
}
