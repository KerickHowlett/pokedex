package mapbcommands

import (
	c "command"
	mf "maps/fetchlocations"
	ml "maps/location"
	qf "query/fetch"
	qs "query/state"
)

type MapBCommand struct {
	state *qs.QueryState[ml.Location]
}

func (m *MapBCommand) Execute() error {
	return mf.FetchLocations(m.state.PreviousURL, m.state, qf.QueryFetch[ml.Location])
}

func (m MapBCommand) GetDescription() string {
	return "Show the previous 20 Pokemon world map locations."
}

func (m MapBCommand) GetName() string {
	return "mapb"
}

func (m *MapBCommand) PrintHelp() {
	c.PrintHelp(m)
}
