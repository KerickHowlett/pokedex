package map_commands_helpers

import (
	"testing"

	ml "maps/location"
	qs "query/state"
	"test_tools/utils"
)

func TestInitializeStateIfNeeded(t *testing.T) {
	t.Run("should return the same state that was provided without mutating it.", func(t *testing.T) {
		arguedState := qs.NewQueryState[ml.Location]()
		returnedState := InitializeStateIfNeeded(arguedState)
		utils.ExpectSameEntity(t, arguedState, returnedState, "state")
	})

	t.Run("should create a new state if none is provided.", func(t *testing.T) {
		if state := InitializeStateIfNeeded(); state == nil {
			t.Error("Expected state to be initialized, but it was nil.")
		}
	})
}
