package pokemon

import (
	ps "pokemon_stat"
	pt "pokemon_type"
	"testing"

	f "test_tools/fixtures"
)

func TestWithBaseExperience(t *testing.T) {
	setup := func() (pokemon *Pokemon, baseExperience int) {
		pokemon = &Pokemon{}
		baseExperience = 100
		WithBaseExperience(baseExperience)(pokemon)
		return pokemon, baseExperience
	}

	t.Run("should assign provided baseExperience to Pokemon.BaseExperience field", func(t *testing.T) {
		if pokemon, baseExperience := setup(); pokemon.BaseExperience != baseExperience {
			t.Errorf("Expected Pokemon.BaseExperience to be %d, but got %d", baseExperience, pokemon.BaseExperience)
		}
	})
}

func TestWithHeight(t *testing.T) {
	setup := func() (pokemon *Pokemon, height int) {
		pokemon = &Pokemon{}
		height = 100
		WithHeight(height)(pokemon)
		return pokemon, height
	}

	t.Run("should assign provided height to Pokemon.Height field", func(t *testing.T) {
		if pokemon, height := setup(); pokemon.Height != height {
			t.Errorf("Expected Pokemon.Height to be %d, but got %d", height, pokemon.Height)
		}
	})
}

func TestWithName(t *testing.T) {
	runWithNameTest := func() (pokemon *Pokemon, name string) {
		pokemon = &Pokemon{}
		name = f.PokemonName
		WithName(name)(pokemon)
		return pokemon, name
	}

	t.Run("should assign provided name to Pokemon.Name field", func(t *testing.T) {
		if pokemon, name := runWithNameTest(); pokemon.Name != name {
			t.Errorf("Expected Pokemon.Name to be %q, but got %q", name, pokemon.Name)
		}
	})
}

func TestWithStat(t *testing.T) {
	setup := func() (pokemon *Pokemon, stat *ps.PokemonStat) {
		pokemon = &Pokemon{Stats: []*ps.PokemonStat{}}
		stat = &ps.PokemonStat{}
		WithStat(stat)(pokemon)
		return pokemon, stat
	}

	t.Run("should assign provided stats to Pokemon.Stats field", func(t *testing.T) {
		if pokemon, stat := setup(); len(pokemon.Stats) != 1 || pokemon.Stats[0] != stat {
			t.Errorf("Expected Pokemon.Stats to be %v, but got %v", stat, pokemon.Stats[0])
		}
	})
}

func TestWithType(t *testing.T) {
	setup := func() (pokemon *Pokemon, pokemonType *pt.PokemonType) {
		pokemon = &Pokemon{Types: []*pt.PokemonType{}}
		pokemonType = &pt.PokemonType{}
		WithType(pokemonType)(pokemon)
		return pokemon, pokemonType
	}

	t.Run("should assign provided type to Pokemon.Types field", func(t *testing.T) {
		if pokemon, type_ := setup(); len(pokemon.Types) != 1 || pokemon.Types[0] != type_ {
			t.Errorf("Expected Pokemon.Types to be %v, but got %v", type_, pokemon.Types[0])
		}
	})
}
