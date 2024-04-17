package pokemon_stat

import (
	"testing"

	el "entity_link"
)

func TestWithBaseStat(t *testing.T) {
	runWithBaseStat := func() (pokemonStat *PokemonStat, baseStat int) {
		baseStat = 100
		pokemonStat = &PokemonStat{}
		WithBaseStat(100)(pokemonStat)
		return pokemonStat, baseStat
	}

	t.Run("should set the base stat value.", func(t *testing.T) {
		if pokemonStat, expected := runWithBaseStat(); pokemonStat.BaseStat != expected {
			t.Errorf("Expected the base stat to be set to %v, but got %v.", expected, pokemonStat.BaseStat)
		}
	})
}

func TestWithStat(t *testing.T) {
	runWithStatTest := func() (pokemonStat *PokemonStat, stat *el.EntityLink) {
		stat = el.NewEntityLink()
		pokemonStat = &PokemonStat{}
		WithStat(stat)(pokemonStat)
		return pokemonStat, stat
	}

	t.Run("should set the stat name and link.", func(t *testing.T) {
		if pokemonStat, expected := runWithStatTest(); pokemonStat.Stat != expected {
			t.Errorf("Expected the stat to be set to %v, but got %v.", expected, pokemonStat.Stat)
		}
	})
}
