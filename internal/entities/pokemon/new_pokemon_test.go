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
}
