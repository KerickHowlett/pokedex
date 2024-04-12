package pokemon

import (
	"testing"

	f "test_tools/fixtures"
)

func TestNewPokemon(t *testing.T) {
	t.Run("should create struct successfully with required Name field set.", func(t *testing.T) {
		pokemon := NewPokemon(WithName(f.StarterTown))

		if pokemon.Name != f.StarterTown {
			t.Errorf("Expected Pokemon.Name to be %q, but got %q", f.StarterTown, pokemon.Name)
		}
	})

	t.Run("should panic if no required Name value is provided.", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected NewPokemon to panic with empty pokemon name")
			}
		}()

		NewPokemon()
	})
}
