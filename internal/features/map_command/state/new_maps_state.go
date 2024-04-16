package maps_state

import l "location"

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
// state := NewMapsState()
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
