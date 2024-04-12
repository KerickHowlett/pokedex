package mapb_commands

import (
	"testing"

	ms "maps/state"
	"test_tools/utils"
)

func TestNewMapBCommand(t *testing.T) {
	t.Run("should create a new MapBCommand with the correct state.", func(t *testing.T) {
		state := ms.NewMapsState()
		command := NewMapBCommand(state)
		utils.ExpectSameEntity(t, state, command.state, "state")
	})

	t.Run("should create a new MapBCommand with a new state if none is provided.", func(t *testing.T) {
		if command := NewMapBCommand(); command.state == nil {
			t.Error("Expected state to be initialized, but it was nil.")
		}
	})
}
