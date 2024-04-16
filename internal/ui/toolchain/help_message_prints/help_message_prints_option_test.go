package help_message_prints

import "testing"

func TestWithCommandsSubTitle(t *testing.T) {
	runWithCommandsSubTitleTest := func() (hmp *HelpMessagePrints, subTitle string) {
		subTitle = "Commands"
		hmp = &HelpMessagePrints{}
		WithCommandsSubTitle(subTitle)(hmp)
		return hmp, subTitle
	}
	t.Run("should set the commands subtitle for the help message.", func(t *testing.T) {
		if hmp, subTitle := runWithCommandsSubTitleTest(); hmp.CommandsSubTitle != subTitle {
			t.Errorf("Expected commands subtitle to be %q but got %q", subTitle, hmp.CommandsSubTitle)
		}
	})
}

func TestWithHorizontalRule(t *testing.T) {
	runWithHorizontalRuleTest := func() (hmp *HelpMessagePrints, hr string) {
		hr = "----------------"
		hmp = &HelpMessagePrints{}
		WithHorizontalRule(hr)(hmp)
		return hmp, hr
	}
	t.Run("should set the horizontal rule for the help message.", func(t *testing.T) {
		if hmp, hr := runWithHorizontalRuleTest(); hmp.HR != hr {
			t.Errorf("Expected horizontal rule to be %q but got %q", hr, hmp.HR)
		}
	})
}

func TestWithTitle(t *testing.T) {
	runWithTitleTest := func() (hmp *HelpMessagePrints, title string) {
		title = "Pokedex Commands"
		hmp = &HelpMessagePrints{}
		WithTitle(title)(hmp)
		return hmp, title
	}
	t.Run("should set the title for the help message.", func(t *testing.T) {
		if hmp, title := runWithTitleTest(); hmp.Title != title {
			t.Errorf("Expected title to be %q but got %q", title, hmp.Title)
		}
	})
}
