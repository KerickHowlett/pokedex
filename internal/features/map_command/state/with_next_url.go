package maps_state

// WithNextURL sets the next URL for the query state.
//
// It is a MapsStateOption function that takes a nullable URL pointer and returns
// a function that sets the NextURL field of the MapsState.
//
// Parameters:
//   - url: The URL to set as the next URL.
//
// Returns:
//   - A function that sets the NextURL field of the MapsState.
//
// Example:
//
//	state := NewMapsState(
//		WithNextURL("https://example.com/query?after=123"),
//	)
func WithNextURL(url *string) MapsStateOption {
	return func(m *MapsState) {
		m.NextURL = url
	}
}
