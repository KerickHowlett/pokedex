package pokemon_encounter

import (
	"testing"

	p "entities/pokemon"
	f "test_tools/fixtures"
)

func TestNewPokemonEncounter(t *testing.T) {
	t.Run("should create a new PokemonEncounter with default values", func(t *testing.T) {
		t.Parallel()
		if NewPokemonEncounter() == nil {
			t.Error("Expected to receive a defined instance of PokemonEncounter, but instead received nil")
		}
	})

	t.Run("should create a new PokemonEncounter with custom options", func(t *testing.T) {
		t.Parallel()
		pokemon := p.NewPokemon(p.WithName(f.PokemonName))
		if pe := NewPokemonEncounter(WithPokemon(pokemon)); pe.Pokemon != pokemon {
			t.Errorf("Expected to receive a PokemonEncounter with Pokemon %v, but instead received %v", pokemon, pe.Pokemon)
		}
	})
}
