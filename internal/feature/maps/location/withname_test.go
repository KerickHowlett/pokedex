package location

import "testing"

func TestWithName(t *testing.T) {

	setup := func() (location *Location, name string) {
		location = &Location{}
		name = "Test Location"

		return location, name

	}

	t.Run("should assign argued name to location.Name field", func(t *testing.T) {
		location, name := setup()

		WithName(name)(location)

		if location.Name != name {
			t.Errorf("Expected location.Name to be %q, but got %q", name, location.Name)
		}
	})
}
