package help_message_prints

import "testing"

func TestNewHelpMessagePrints(t *testing.T) {
	runNewHelpMessagePrintsTest := func() (hmp *HelpMessagePrints, commandSubTitle string, hr string, title string) {
		commandSubTitle = "Test Subtitle"
		hr = "*****************************"
		title = "Test Title"
		hmp = NewHelpMessagePrints(
			WithCommandsSubTitle(commandSubTitle),
			WithHorizontalRule(hr),
			WithTitle(title),
		)
		return hmp, commandSubTitle, hr, title
	}
	t.Run("should create a new HelpMessagePrints instance with the provided 'commandSubTitle' option.", func(t *testing.T) {
		t.Parallel()
		if hmp, commandSubTitle, _, _ := runNewHelpMessagePrintsTest(); hmp.CommandsSubTitle != commandSubTitle {
			t.Errorf("Expected commands subtitle to be %q but got %q", commandSubTitle, hmp.CommandsSubTitle)
		}
	})

	t.Run("should create a new HelpMessagePrints instance with the provided 'hr' option.", func(t *testing.T) {
		t.Parallel()
		if hmp, _, hr, _ := runNewHelpMessagePrintsTest(); hmp.HR != hr {
			t.Errorf("Expected horizontal rule to be %q but got %q", hr, hmp.HR)
		}
	})

	t.Run("should create a new HelpMessagePrints instance with the provided 'title' option.", func(t *testing.T) {
		t.Parallel()
		if hmp, _, _, title := runNewHelpMessagePrintsTest(); hmp.Title != title {
			t.Errorf("Expected title to be %q but got %q", title, hmp.Title)
		}
	})
}
