package location_area

import p "explore/pokemon"

// LocationArea houses the essential information of a given location area.
//
// Fields:
//   - Encounters: The Pokemon that can be encountered in the location area.
type LocationArea struct {
	// Encounters represents the Pokemon that can be encountered in the location area.
	Encounters []p.Pokemon `json:"pokemon_encounters"`
}
