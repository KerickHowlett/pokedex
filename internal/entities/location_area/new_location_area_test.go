package location_area

import (
	el "entity_link"
	"testing"

	pe "pokemon_encounter"
	"test_tools/fixtures"
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
		pokemon := el.NewEntityLink(el.WithName(fixtures.PokemonName))
		encounter := pe.NewPokemonEncounter(pe.WithPokemon(pokemon))
		locationArea := NewLocationArea(WithPokemonEncounter(*encounter))

		if actualEncounter := locationArea.Encounters[0]; actualEncounter != *encounter {
			t.Errorf("Expected encounter %v, got %v", *pokemon, actualEncounter)
		}
	})
}
