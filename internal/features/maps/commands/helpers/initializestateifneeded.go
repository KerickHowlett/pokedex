package mapcommandshelpers

import (
	ml "maps/location"
	qs "query/state"
)

func InitializeStateIfNeeded(state ...*qs.QueryState[ml.Location]) *qs.QueryState[ml.Location] {
	if state != nil && state[0] != nil {
		return state[0]
	}
	return qs.NewQueryState[ml.Location]()
}
