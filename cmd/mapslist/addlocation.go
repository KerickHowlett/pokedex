package mapslist

import l "github.com/KerickHowlett/pokedexcli/cmd/location"

func AddLocation(location l.Location) MapsListOption {
	return func(m *MapsList) {
		m.Locations = append(m.Locations, location)
	}
}
