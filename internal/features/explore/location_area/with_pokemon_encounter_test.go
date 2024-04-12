package location_area

import (
	"testing"

	p "explore/pokemon"
)

func TestWithPokemonEncounter(t *testing.T) {
	runWithPokemonEncounterTest := func() (encounter p.Pokemon, locationArea LocationArea) {
		encounter, locationArea = p.Pokemon{}, LocationArea{}
		WithPokemonEncounter(encounter)(&locationArea)
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
		if encounter, locationArea := runWithPokemonEncounterTest(); locationArea.Encounters[0] != encounter {
			t.Errorf("Expected encounter %v, got %v", encounter, locationArea.Encounters[0])
		}

	})
}
