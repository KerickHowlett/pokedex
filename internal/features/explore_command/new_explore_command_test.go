package explore_command

import (
	"testing"

	f "test_tools/fixtures"
	"test_tools/utils"
)

func TestNewExploreCommand(t *testing.T) {
	t.Run("should create a new ExploreCommand instance.", func(t *testing.T) {
		t.Parallel()
		if NewExploreCommand(WithAPIEndpoint(f.APIEndpoint)) == nil {
			t.Error("Expected a new ExploreCommand instance, but got nil.")
		}
	})

	t.Run("should panic if API endpoint is not provided.", func(t *testing.T) {
		utils.ExpectPanic(t, NewExploreCommand)
	})
}
