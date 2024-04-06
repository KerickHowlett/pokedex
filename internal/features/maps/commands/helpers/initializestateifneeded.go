package mapcommandshelpers

import (
	ml "maps/location"
	qs "query/state"
)

// InitializeStateIfNeeded initializes the query state if it is needed.
//
// If the state is not nil and the first element is not nil, it returns the first element.
// Otherwise, it returns a new query state.
//
// Parameters:
//   - state: The query state to initialize. Nullable.
//
// Returns:
//   - The query state if it is not nil, otherwise a new query state.
//
// Example:
//
//	 arguedState := qs.NewQueryState[ml.Location]()
//		state1 := InitializeStateIfNeeded(arguedState)
//		state2 := InitializeStateIfNeeded()
func InitializeStateIfNeeded(state ...*qs.QueryState[ml.Location]) *qs.QueryState[ml.Location] {
	if state != nil && state[0] != nil {
		return state[0]
	}
	return qs.NewQueryState[ml.Location]()
}
