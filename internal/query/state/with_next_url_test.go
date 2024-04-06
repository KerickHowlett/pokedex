package querystate

import (
	"testing"

	f "test_tools/fixtures"
)

func TestWithNextURL(t *testing.T) {
	var nilString *string = nil
	runWithNextURLTest := func(url *string, defaultNextURL ...*string) *queryTestState {
		state := &queryTestState{}
		if len(defaultNextURL) > 0 && defaultNextURL[0] != nil {
			state.NextURL = defaultNextURL[0]
		}

		WithNextURL[result](url)(state)

		return state
	}

	t.Run("should set QueryState[TResult].NextURL with the provided URL.", func(t *testing.T) {
		if state := runWithNextURLTest(&f.APIEndpoint); *state.NextURL != f.APIEndpoint {
			t.Errorf("Expected QueryState[TResult].NextURL to be %q, but got %q", f.APIEndpoint, *state.NextURL)
		}
	})

	t.Run("should set QueryState[TResult].NextURL with nil URL.", func(t *testing.T) {
		if state := runWithNextURLTest(nilString, &f.APIEndpoint); state.NextURL != nilString {
			t.Errorf("Expected QueryState[TResult].NextURL to be nil, but got non-nil")
		}
	})
}
