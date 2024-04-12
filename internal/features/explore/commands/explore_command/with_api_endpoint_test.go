package explore

import (
	"testing"

	f "test_tools/fixtures"
)

func TestWithAPIEndpoint(t *testing.T) {
	runWithAPIEndpointTests := func() *ExploreCommand {
		command := &ExploreCommand{}
		WithAPIEndpoint(f.APIEndpoint)(command)
		return command
	}

	t.Run("should set the API endpoint for the ExploreCommand instance.", func(t *testing.T) {
		if command := runWithAPIEndpointTests(); command.apiEndpoint != f.APIEndpoint {
			t.Errorf("Expected apiEndpoint to be %s, but got %s", f.APIEndpoint, command.apiEndpoint)
		}
	})
}
