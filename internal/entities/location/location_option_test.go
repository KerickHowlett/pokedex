package location

import (
	"testing"

	f "test_tools/fixtures"
)

func TestWithName(t *testing.T) {
	setup := func() (location *Location, name string) {
		location = &Location{}
		name = f.StarterTown
		return location, name
	}

	t.Run("should assign argued name to Location.Name field", func(t *testing.T) {
		location, name := setup()

		WithName(name)(location)

		if location.Name != name {
			t.Errorf("Expected Location.Name to be %q, but got %q", name, location.Name)
		}
	})
}
