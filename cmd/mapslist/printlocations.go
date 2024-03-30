package mapslist

import "fmt"

func (m *MapsList) PrintLocations() error {
	if len(m.Locations) == 0 {
		return fmt.Errorf("no maps found")
	}

	for _, locations := range m.Locations {
		locationName := locations.Name
		fmt.Println(locationName)
	}

	return nil
}
