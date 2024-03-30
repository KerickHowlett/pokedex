package mapslist

func WithNextURL(url *string) MapsListOption {
	return func(m *MapsList) {
		m.NextURL = url
	}
}
