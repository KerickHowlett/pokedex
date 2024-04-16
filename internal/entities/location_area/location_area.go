package location_area

import pe "pokemon_encounter"

// LocationArea houses the essential information of a given location area.
//
// Fields:
//   - Encounters: The Pokemon that can be encountered in the location area.
type LocationArea struct {
	// Encounters represents the Pokemon that can be encountered in the location area.
	Encounters []pe.PokemonEncounter `json:"pokemon_encounters"`
}
