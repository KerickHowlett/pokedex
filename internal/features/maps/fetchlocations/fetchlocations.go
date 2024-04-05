package fetchlocations

import (
	"fmt"

	l "maps/location"
	qs "query/state"
)

type MapsQueryState = qs.QueryState[l.Location]
type QueryFetchFunc = func(url string, queryState *MapsQueryState) error

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

	for _, location := range state.Results {
		fmt.Println(location.Name)
	}

	return nil
}
