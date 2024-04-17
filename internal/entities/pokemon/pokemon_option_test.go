package pokemon

import (
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

	t.Run("should assign argued baseExperience to Pokemon.BaseExperience field", func(t *testing.T) {
		if pokemon, baseExperience := setup(); pokemon.BaseExperience != baseExperience {
			t.Errorf("Expected Pokemon.BaseExperience to be %d, but got %d", baseExperience, pokemon.BaseExperience)
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

	t.Run("should assign argued name to Pokemon.Name field", func(t *testing.T) {
		if pokemon, name := runWithNameTest(); pokemon.Name != name {
			t.Errorf("Expected Pokemon.Name to be %q, but got %q", name, pokemon.Name)
		}
	})
}
