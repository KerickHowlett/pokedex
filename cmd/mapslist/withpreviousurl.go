package mapslist

func WithPreviousURL(url *string) MapsListOption {
	return func(m *MapsList) {
		m.PreviousURL = url
	}
}
