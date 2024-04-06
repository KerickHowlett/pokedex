package location

import (
	"testing"

	f "test_tools/fixtures"
)

func TestNewLocation(t *testing.T) {
	t.Run("should create struct successfully with required Name field set.", func(t *testing.T) {
		location := NewLocation(WithName(f.StarterTown))

		if location.Name != f.StarterTown {
			t.Errorf("Expected location.name to be %q, but got %q", f.StarterTown, location.Name)
		}
	})

	t.Run("should panic if no required Name value is provided.", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected NewLocation to panic with empty location name")
			}
		}()

		NewLocation()
	})
}
