package maps_state

import l "entities/location"

// NewMapsState creates a new instance of MapsState with the specified options.
//
// Parameters:
//   - options: The options to configure the MapsState.
//
// Returns:
//   - A new instance of MapsState.
//
// Example:
//
//		state := NewMapsState(
//			WithNextURL("https://example.com/query?after=123"),
//			WithPreviousURL("https://example.com/query?before=123"),
//	     WithLocations("result1"),
//	     WithLocations("result2"),
//
// )
func NewMapsState(options ...MapsStateOption) (mapsState *MapsState) {
	mapsState = &MapsState{
		NextURL:     nil,
		PreviousURL: nil,
		Locations:   []l.Location{},
	}

	for _, option := range options {
		option(mapsState)
	}

	return mapsState
}
