package map_commands

import (
	"testing"

	ml "maps/location"
	qs "query/state"
	"test_tools/utils"
)

func TestNewMapCommand(t *testing.T) {
	t.Run("should create a new MapCommand with the correct state.", func(t *testing.T) {
		state := qs.NewQueryState[ml.Location]()
		command := NewMapCommand(state)
		utils.ExpectSameEntity(t, state, command.state, "state")
	})

	t.Run("should create a new MapCommand with a new state if none is provided.", func(t *testing.T) {
		if command := NewMapCommand(); command.state == nil {
			t.Error("Expected state to be initialized, but it was nil.")
		}
	})
}
