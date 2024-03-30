package mapslist

import l "github.com/KerickHowlett/pokedexcli/cmd/location"

type MapsList struct {
	NextURL     *string      `json:"next"`
	PreviousURL *string      `json:"previous"`
	Locations   []l.Location `json:"results"`
}
