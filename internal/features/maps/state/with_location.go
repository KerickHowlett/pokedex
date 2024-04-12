package maps_state

import l "maps/location"

// WithLocation is a MapsStateOption function that appends the given result to
// the query state's Results slice.
//
// Parameters:
//   - result: The result to append to the query state's Results slice.
//
// Returns:
//   - A function that appends the given result to the query state's Results slice.
//
// Example:
//
//	state := NewMapsState(
//		WithLocation("result1"),
//		WithLocation("result2"),
//	)
func WithLocation(location l.Location) MapsStateOption {
	return func(m *MapsState) {
		m.Locations = append(m.Locations, location)
	}
}
