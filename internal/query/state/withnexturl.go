package querystate

func WithNextURL[TResult any](url *string) QueryStateOption[TResult] {
	return func(m *QueryState[TResult]) {
		m.NextURL = url
	}
}
