package mapcommands

import (
	c "command"
	mf "maps/fetchlocations"
	ml "maps/location"
	qf "query/fetch"
	qs "query/state"
)

type MapCommand struct {
	state *qs.QueryState[ml.Location]
}

func (m *MapCommand) Execute() error {
	return mf.FetchLocations(m.state.NextURL, m.state, qf.QueryFetch[ml.Location])
}

func (m MapCommand) GetDescription() string {
	return "Show the next 20 Pokemon world map locations."
}

func (m MapCommand) GetName() string {
	return "map"
}

func (m *MapCommand) PrintHelp() {
	c.PrintHelp(m)
}
