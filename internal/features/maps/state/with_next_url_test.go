package maps_state

import (
	"testing"

	f "test_tools/fixtures"
)

func TestWithNextURL(t *testing.T) {
	var nilString *string = nil
	runWithNextURLTest := func(url *string, defaultNextURL ...*string) *MapsState {
		state := &MapsState{}
		if len(defaultNextURL) > 0 && defaultNextURL[0] != nil {
			state.NextURL = defaultNextURL[0]
		}

		WithNextURL(url)(state)

		return state
	}

	t.Run("should set MapsState.NextURL with the provided URL.", func(t *testing.T) {
		if state := runWithNextURLTest(&f.APIEndpoint); *state.NextURL != f.APIEndpoint {
			t.Errorf("Expected MapsState.NextURL to be %q, but got %q", f.APIEndpoint, *state.NextURL)
		}
	})

	t.Run("should set MapsState.NextURL with nil URL.", func(t *testing.T) {
		if state := runWithNextURLTest(nilString, &f.APIEndpoint); state.NextURL != nilString {
			t.Errorf("Expected MapsState.NextURL to be nil, but got non-nil")
		}
	})
}
