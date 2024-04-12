package pokemon

import (
	"testing"

	f "test_tools/fixtures"
)

func TestWithName(t *testing.T) {
	setup := func() (pokemon *Pokemon, name string) {
		pokemon = &Pokemon{}
		name = f.PokemonName

		return pokemon, name

	}

	t.Run("should assign argued name to Pokemon.Name field", func(t *testing.T) {
		pokemon, name := setup()

		WithName(name)(pokemon)

		if pokemon.Name != name {
			t.Errorf("Expected Pokemon.Name to be %q, but got %q", name, pokemon.Name)
		}
	})
}
