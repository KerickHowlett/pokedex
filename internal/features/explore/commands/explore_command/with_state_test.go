package explore

import (
	"testing"

	la "explore/location_area"
)

func TestWithState(t *testing.T) {
	runWithStateTest := func() (state *la.LocationArea, command *ExploreCommand) {
		state, command = &la.LocationArea{}, &ExploreCommand{}
		WithState(state)(command)
		return state, command
	}

	t.Run("should set the state for the ExploreCommand instance.", func(t *testing.T) {
		state, command := runWithStateTest()
		if command.state != state {
			t.Errorf("Expected state to be %v, but got %v", state, command.state)
		}
	})
}
