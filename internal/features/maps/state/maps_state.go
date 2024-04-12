package maps_state

import l "maps/location"

// MapsState represents the state of a query result.
//
// Generic Constraints:
//
//	T - The query result's type, whether it'd be a primitive type or a struct.
//
// Fields:
//   - NextURL: Represents the URL for the next page of results. Nullable.
//   - PreviousURL: Represents the URL of the previous page of results. Nullable.
//   - Locations: Represents the results of the query, which are all locations
//     that can be visited within the various Pokemon games.
//
// Example:
//
//	state := &MapsState{}
type MapsState struct {
	NextURL     *string      `json:"next"`
	PreviousURL *string      `json:"previous"`
	Locations   []l.Location `json:"results"`
}
