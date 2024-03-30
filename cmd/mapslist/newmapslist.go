package mapslist

import l "github.com/KerickHowlett/pokedexcli/cmd/location"

func NewMapsList(options ...MapsListOption) *MapsList {
	maps := &MapsList{
		NextURL:     nil,
		PreviousURL: nil,
		Locations:   []l.Location{},
	}

	for _, option := range options {
		option(maps)
	}

	return maps
}
