package fetch_locations

import (
	"fmt"

	l "maps/location"
	qs "query/state"
)

type MapsQueryState = qs.QueryState[l.Location]
type QueryFetchFunc = func(url string, queryState *MapsQueryState) error

// FetchLocations fetches locations from a given URL -- state.NextURL or
// state.PreviousURL -- using the provided fetch function and updates the query state.
// It prints the names of the fetched locations and returns an error if any occurred.
//
// Parameters:
//   - url: A pointer to a string containing the URL to fetch locations from.
//   - state: A pointer to a QueryState struct instance.
//   - fetch: A function that fetches locations from a given URL.
//
// Returns:
//   - An error if any occurred.
//
// Example usage:
//
//	err := FetchLocations(&state.NextURL, state, fetch)
func FetchLocations(url *string, state *MapsQueryState, fetch QueryFetchFunc) error {
	if fetch == nil {
		panic("fetch function is required")
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

	if len(state.Results) == 0 {
		return fmt.Errorf("no maps were found")
	}

	fmt.Println("Pokemon Maps:")
	for _, location := range state.Results {
		fmt.Println(location.Name)
	}

	return nil
}
