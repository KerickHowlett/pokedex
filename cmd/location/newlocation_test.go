package location

import "testing"

func TestNewLocation(t *testing.T) {
	t.Run("Valid Location Name", func(t *testing.T) {
		const starterTown = "Pallet Town"
		location := NewLocation(WithName(starterTown))

		if location.Name != starterTown {
			t.Errorf("Expected location.name to be %q, but got %q", starterTown, location.Name)
		}
	})

	t.Run("Empty Location Name", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected NewLocation to panic with empty location name")
			}
		}()

		NewLocation()
	})
}
