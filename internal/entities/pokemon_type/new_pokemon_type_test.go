package pokemon_type

import (
	"testing"
)

func TestNewPokemonType(t *testing.T) {
	t.Run("should create a new PokemonType.", func(t *testing.T) {
		t.Parallel()
		if pokemonType := NewPokemonType(); pokemonType == nil {
			t.Error("Expected a new PokemonType to be instantiated, but instead, its nil.")
		}
	})

	t.Run("should have a non-nil value set to 'type'.", func(t *testing.T) {
		t.Parallel()
		if pokemonType := NewPokemonType(); pokemonType.Type == nil {
			t.Error("Expected a new PokemonType to be instantiated, but instead, its nil.")
		}
	})
}
