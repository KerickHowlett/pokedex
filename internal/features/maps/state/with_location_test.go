package maps_state

import (
	"testing"

	l "maps/location"
	f "test_tools/fixtures"
)

func TestWithLocation(t *testing.T) {
	runWithLocationTest := func() (location *l.Location, state *MapsState) {
		location = l.NewLocation(l.WithName(f.StarterTown))
		state = &MapsState{}

		WithLocation(*location)(state)

		return location, state
	}

	t.Run("should add location to the MapsState.Results field.", func(t *testing.T) {
		if _, state := runWithLocationTest(); len(state.Locations) != 1 {
			t.Errorf("Expected MapsState.NextURL to have 1 item, but got %d", len(state.Locations))
		}
	})

	t.Run("should add Location with unchanged values.", func(t *testing.T) {
		if item, state := runWithLocationTest(); state.Locations[0] != *item {
			t.Errorf("Expected MapsState.NextURL[0] to be %v, but got %v", item, state.Locations[0])
		}
	})
}
