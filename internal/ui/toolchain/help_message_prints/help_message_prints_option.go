package help_message_prints

type HelpMessagePrintsOption func(*HelpMessagePrints)

// WithCommandsSubTitle sets the commands subtitle for the help message.
//
// Parameters:
//   - subTitle: The subtitle for the commands section.
//
// Returns:
//   - HelpMessagePrintsOption: The option to set the commands subtitle.
//
// Example usage:
//
// hmp := NewHelpMessagePrints(WithCommandsSubTitle("Commands SubTitle"))
func WithCommandsSubTitle(subTitle string) HelpMessagePrintsOption {
	return func(h *HelpMessagePrints) {
		h.CommandsSubTitle = subTitle
	}
}

// WithHorizontalRule sets the horizontal rule for the help message.
//
// Parameters:
//   - hr: The horizontal rule.
//
// Returns:
//   - HelpMessagePrintsOption: The option to set the horizontal rule.
//
// Example usage:
//
// hmp := NewHelpMessagePrints(WithHorizontalRule("***************"))
func WithHorizontalRule(hr string) HelpMessagePrintsOption {
	return func(h *HelpMessagePrints) {
		h.HR = hr
	}
}

// WithTitle sets the title for the help message.
//
// Parameters:
//   - title: The title for the help message.
//
// Returns:
//   - HelpMessagePrintsOption: The option to set the title.
//
// Example usage:
//
// hmp := NewHelpMessagePrints(WithTitle("Help Title"))
func WithTitle(title string) HelpMessagePrintsOption {
	return func(h *HelpMessagePrints) {
		h.Title = title
	}
}
