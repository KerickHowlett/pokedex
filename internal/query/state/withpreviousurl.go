package querystate

func WithPreviousURL[TResult any](url *string) QueryStateOption[TResult] {
	return func(m *QueryState[TResult]) {
		m.PreviousURL = url
	}
}
