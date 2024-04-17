package pokemon_stat

import "testing"

func TestNewPokemonStat(t *testing.T) {
	t.Run("should create a new PokemonStat.", func(t *testing.T) {
		t.Parallel()
		if pokemonStat := NewPokemonStat(); pokemonStat == nil {
			t.Error("Expected a new PokemonStat to be instantiated, but instead, its nil.")
		}
	})

	t.Run("should have a non-nil value set to 'type'.", func(t *testing.T) {
		t.Parallel()
		if pokemonStat := NewPokemonStat(); pokemonStat.Stat == nil {
			t.Error("Expected a new PokemonStat to be instantiated, but instead, its nil.")
		}
	})
}
