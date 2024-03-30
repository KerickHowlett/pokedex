package location

import "testing"

func TestWithName(t *testing.T) {
	location := &Location{}
	name := "Test Location"

	WithName(name)(location)

	if location.Name != name {
		t.Errorf("Expected location.name to be %q, but got %q", name, location.Name)
	}
}
