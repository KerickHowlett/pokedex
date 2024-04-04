package querystate

func WithResult[TResult any](result TResult) QueryStateOption[TResult] {
	return func(m *QueryState[TResult]) {
		m.Results = append(m.Results, result)
	}
}
