package fetch_encounters

import (
	"fmt"

	la "explore/location_area"
	qf "query/fetch"
)

// FetchEncounters fetches Pokemon that can be encountered within the queried
// LocationArea using the provided fetch function and updates the query state.
//
// Parameters:
//   - url: A string containing the URL to fetch locations from.
//   - state: A pointer to a LocationArea struct instance.
//   - fetch: A function that fetches locations from a given URL.
//
// Returns:
//   - An error if any occurred.
//
// Example usage:
//
//	err := FetchLocations(&state.NextURL, state, fetch)
func FetchEncounters(url string, state *la.LocationArea, fetchFunc ...qf.QueryFetchFunc[la.LocationArea]) error {
	fetch := qf.QueryFetch[la.LocationArea]
	if len(fetchFunc) > 0 {
		fetch = fetchFunc[0]
	}

	if state == nil {
		return fmt.Errorf("a query state is required to fetch locations")
	}

	if err := fetch(url, state); err != nil {
		return err
	}

	if len(state.Encounters) == 0 {
		return fmt.Errorf("no pokemon were found in the area")
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range state.Encounters {
		fmt.Printf("  - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
