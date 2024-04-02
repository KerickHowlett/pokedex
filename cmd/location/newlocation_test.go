package location

import "testing"

func TestNewLocation(t *testing.T) {
	const starterTown = "Pallet Town"

	t.Run("should create struct successfully with required Name field set.", func(t *testing.T) {
		location := NewLocation(WithName(starterTown))

		if location.Name != starterTown {
			t.Errorf("Expected location.name to be %q, but got %q", starterTown, location.Name)
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
