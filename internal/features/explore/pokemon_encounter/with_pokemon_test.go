package pokemon_encounter

import (
	"testing"

	p "explore/pokemon"
	f "test_tools/fixtures"
)

func TestWithPokemon(t *testing.T) {
	runWithPokemonTest := func() (pe *PokemonEncounter, pokemon *p.Pokemon) {
		pe = &PokemonEncounter{}
		pokemon = p.NewPokemon(p.WithName(f.PokemonName))
		WithPokemon(pokemon)(pe)
		return pe, pokemon
	}

	t.Run("should set provided Pokemon to the PokemonEncounter struct", func(t *testing.T) {
		pe, expectedPokemon := runWithPokemonTest()
		if pe.Pokemon != expectedPokemon {
			t.Error("Expected Pokemon to be set to the PokemonEncounter struct, but it was not")
		}
	})
}
