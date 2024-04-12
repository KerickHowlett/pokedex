package maps_state

import "testing"

func TestNewMapsState(t *testing.T) {
	expectURLToBeNil := func(url *string, label string) {
		if url != nil {
			t.Errorf("Expected %s to be nil, but got non-nil", label)
		}
	}

	t.Run("should create MapsState with it's NextURL field set to nil by default.", func(t *testing.T) {
		state := NewMapsState()
		expectURLToBeNil(state.NextURL, "MapsState.NextURL")
	})

	t.Run("should create MapsState with it's PreviousURL field set to nil by default.", func(t *testing.T) {
		state := NewMapsState()
		expectURLToBeNil(state.PreviousURL, "MapsState.PreviousURL")
	})

	t.Run("should create MapsState with it's Results field set with an empty slice by default.", func(t *testing.T) {
		state := NewMapsState()
		if totalResults := len(state.Locations); totalResults != 0 {
			t.Errorf("Expected results to be empty, but got %d elements", totalResults)
		}
	})
}
