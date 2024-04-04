package querystate

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
