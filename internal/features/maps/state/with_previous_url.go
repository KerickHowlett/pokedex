package maps_state

// WithPreviousURL sets the next URL for the query state.
//
// It is a MapsStateOption function that takes a nullable URL pointer and returns
// a function that sets the PreviousURL field of the MapsState.
//
// Parameters:
//   - url: The URL to set as the next URL.
//
// Returns:
//   - A function that sets the PreviousURL field of the MapsState.
//
// Example:
//
//	state := NewMapsState(
//		WithPreviousURL("https://example.com/query?after=123"),
//	)
func WithPreviousURL(url *string) MapsStateOption {
	return func(m *MapsState) {
		m.PreviousURL = url
	}
}
