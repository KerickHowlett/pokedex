package map_commands_helpers

import ms "maps/state"

// InitializeStateIfNeeded initializes the maps state if it is needed.
//
// If the state is not nil and the first element is not nil, it returns the first element.
// Otherwise, it returns a new maps state.
//
// Parameters:
//   - state: The maps state to initialize. Nullable.
//
// Returns:
//   - The maps state if it is not nil, otherwise a new maps state.
//
// Example:
//
//	 arguedState := qs.NewMapsState()
//		state1 := InitializeStateIfNeeded(arguedState)
//		state2 := InitializeStateIfNeeded()
func InitializeStateIfNeeded(state ...*ms.MapsState) *ms.MapsState {
	if state != nil && state[0] != nil {
		return state[0]
	}
	return ms.NewMapsState()
}
