package fetch_locations

import (
	"fmt"

	ms "maps/state"
	qf "query/fetch"
)

// FetchLocations fetches locations from a given URL -- state.NextURL or
// state.PreviousURL -- using the provided fetch function and updates the query state.
// It prints the names of the fetched locations and returns an error if any occurred.
//
// Parameters:
//   - url: A pointer to a string containing the URL to fetch locations from.
//   - state: A pointer to a MapsState struct instance.
//   - fetch: A function that fetches locations from a given URL.
//
// Returns:
//   - An error if any occurred.
//
// Example usage:
//
//	err := FetchLocations(&state.NextURL, state, fetch)
func FetchLocations(url *string, state *ms.MapsState, fetchFunc ...qf.QueryFetchFunc[ms.MapsState]) error {
	fetch := qf.QueryFetch[ms.MapsState]
	if len(fetchFunc) > 0 {
		fetch = fetchFunc[0]
	}

	if state == nil {
		return fmt.Errorf("a query state is required to fetch locations")
	}

	if url == nil {
		return fmt.Errorf("no more maps to fetch")
	}

	if err := fetch(*url, state); err != nil {
		return err
	}

	if len(state.Locations) == 0 {
		return fmt.Errorf("no maps were found")
	}

	fmt.Println("Pokemon Maps:")
	for _, location := range state.Locations {
		fmt.Printf("  - %s\n", location.Name)
	}

	return nil
}
