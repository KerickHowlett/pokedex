package toolchain

import (
	"testing"
	hmp "toochain/help_message_prints"

	c "command"
	m "test_tools/mocks/command"
	"test_tools/utils"
)

func TestWithCommand(t *testing.T) {
	runWithCommandTest := func() (toolchain *Toolchain, command *m.MockCommand) {
		toolchain = &Toolchain{commands: &c.Commands{}}
		command = m.NewMockCommand()
		WithCommand(command)(toolchain)
		return toolchain, command
	}

	t.Run("When adding a command with a unique name", func(t *testing.T) {
		t.Parallel()
		t.Run("Then it should add the command to the toolchain.", func(t *testing.T) {
			toolchain, command := runWithCommandTest()
			if _, commandExists := (*toolchain.commands)[command.GetName()]; !commandExists {
				t.Error("Expected command to exist in toolchain")
			}
		})
	})

	t.Run("When adding a command with a name matching a pre-existing command", func(t *testing.T) {
		t.Parallel()
		t.Run("Then it should panic to prevent the command overwrite.", func(t *testing.T) {
			toolchain, command := runWithCommandTest()
			utils.ExpectPanic(t, func() { WithCommand(command)(toolchain) })
		})
	})
}

func TestWithHelpMessagePrints(t *testing.T) {
	runWithHelpMessagePrintsTest := func() (actual *hmp.HelpMessagePrints, expected *hmp.HelpMessagePrints) {
		expected = hmp.NewHelpMessagePrints(
			hmp.WithCommandsSubTitle("SubTitle"),
			hmp.WithTitle("Title"),
			hmp.WithHorizontalRule("****"),
		)
		toolchain := &Toolchain{}
		WithHelpMessagePrints(expected)(toolchain)
		return toolchain.prints, expected
	}

	t.Run("should set the help message prints to the 'prints' field.", func(t *testing.T) {
		toolchain := &Toolchain{}
		prints := hmp.NewHelpMessagePrints()
		WithHelpMessagePrints(prints)(toolchain)
		if actual, expected := runWithHelpMessagePrintsTest(); actual != expected {
			t.Errorf("Expected: %v, but got: %v", expected, actual)
		}
	})
}
