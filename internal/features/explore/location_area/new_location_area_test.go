package location_area

import (
	p "explore/pokemon"
	"test_tools/fixtures"
	"testing"
)

func TestNewLocationArea(t *testing.T) {
	t.Run("should instantiate new struct with empty encounters list", func(t *testing.T) {
		t.Parallel()
		if locationArea := NewLocationArea(); len(locationArea.Encounters) > 0 {
			t.Errorf("Expected new struct instance, got %v", locationArea)
		}
	})

	t.Run("should instantiate new struct with added options", func(t *testing.T) {
		t.Parallel()
		pokemon := p.NewPokemon(p.WithName(fixtures.PokemonName))
		locationArea := NewLocationArea(WithPokemonEncounter(*pokemon))

		if actualEncounter := locationArea.Encounters[0]; actualEncounter != *pokemon {
			t.Errorf("Expected encounter %v, got %v", *pokemon, actualEncounter)
		}
	})
}
