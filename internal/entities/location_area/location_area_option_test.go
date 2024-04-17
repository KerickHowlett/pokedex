package location_area

import (
	"testing"

	el "entity_link"
	pe "pokemon_encounter"
	f "test_tools/fixtures"
)

func TestWithPokemonEncounter(t *testing.T) {
	runWithPokemonEncounterTest := func() (encounter *pe.PokemonEncounter, locationArea LocationArea) {
		locationArea = LocationArea{}

		pokemon := el.NewEntityLink(el.WithName(f.PokemonName))
		encounter = pe.NewPokemonEncounter(pe.WithPokemon(pokemon))

		WithPokemonEncounter(*encounter)(&locationArea)

		return encounter, locationArea
	}

	t.Run("should add a Pokemon encounter to the LocationArea", func(t *testing.T) {
		t.Parallel()
		if _, locationArea := runWithPokemonEncounterTest(); len(locationArea.Encounters) != 1 {
			t.Errorf("Expected 1 encounter, got %d", len(locationArea.Encounters))
		}
	})

	t.Run("should add the correct Pokemon Encounter to the LocationArea", func(t *testing.T) {
		t.Parallel()
		if encounter, locationArea := runWithPokemonEncounterTest(); locationArea.Encounters[0] != *encounter {
			t.Errorf("Expected encounter %v, got %v", *encounter, locationArea.Encounters[0])
		}

	})
}
