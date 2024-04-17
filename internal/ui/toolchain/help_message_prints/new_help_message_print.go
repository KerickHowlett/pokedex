package help_message_prints

// NewHelpMessagePrints creates a new HelpMessagePrints instance.
//
// Parameters:
//   - options: The options to set the HelpMessagePrints fields.
//
// Returns:
//   - *HelpMessagePrints: The new HelpMessagePrints instance.
//
// Example usage:
//
//	hmp := NewHelpMessagePrints(
//	  WithCommandsSubTitle("Commands"),
//	  WithHorizontalRule("----------------"),
//	  WithTitle("Pokedex Commands"),
//	)
func NewHelpMessagePrints(options ...HelpMessagePrintsOption) *HelpMessagePrints {
	hmp := &HelpMessagePrints{
		CommandsSubTitle: "Pokedex Commands:",
		HR:               "----------------",
		Title:            "Welcome to the Pokedex CLI!",
	}
	for _, option := range options {
		option(hmp)
	}
	return hmp
}
