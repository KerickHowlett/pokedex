package explore

import (
	"testing"

	f "test_tools/fixtures"
)

func TestNewExploreCommand(t *testing.T) {
	t.Run("should create a new ExploreCommand instance.", func(t *testing.T) {
		t.Parallel()
		if NewExploreCommand(WithAPIEndpoint(f.APIEndpoint)) == nil {
			t.Error("Expected a new ExploreCommand instance, but got nil.")
		}
	})

	t.Run("should panic if API endpoint is not provided.", func(t *testing.T) {
		t.Parallel()
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected a panic, but got nil.")
			}
		}()
		NewExploreCommand()
	})
}
