package explore

import la "explore/location_area"

// WithState sets the state for the ExploreCommand instance.
//
// Parameters:
//   - state: A pointer to a LocationArea struct instance.
//
// Returns:
//   - A new ExploreCommandOption instance.
//
// Example usage:
//
//	command := NewExploreCommand(WithState(la.NewLocationArea()))
func WithState(state *la.LocationArea) func(*ExploreCommand) {
	return func(e *ExploreCommand) {
		e.state = state
	}
}
