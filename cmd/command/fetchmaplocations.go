package command

import (
	"fmt"

	l "location"
	qf "query/fetch"
	qs "query/state"
)

func fetchMapLocations(url *string, state *qs.QueryState[l.Location]) error {
	if url == nil {
		return fmt.Errorf("no more maps to fetch")
	}

	if err := qf.QueryFetch(*url, state); err != nil {
		return err
	}

	if len(state.Results) == 0 {
		return fmt.Errorf("no maps found")
	}

	for _, location := range state.Results {
		fmt.Println(location.Name)
	}

	return nil
}
