package toolchain

import (
	c "command"
	"fmt"
	"testing"

	f "test_tools/fixtures"
	mc "test_tools/mocks/command"
	"test_tools/utils"
)

func TestPrintHelpMessage(t *testing.T) {
	stdout := utils.NewPrintStorage()

	setup := func() (actualOutput, expectedOutput string) {
		command1 := mc.NewMockCommand(
			mc.WithName("Command 1"),
			mc.WithDescription("Command 1 help message."),
		)
		command2 := mc.NewMockCommand(
			mc.WithName("Command 2"),
			mc.WithDescription("Command 2 help message."),
		)
		toolchain := NewToolchain(
			WithCommand(command1),
			WithCommand(command2),
		)

		actualOutput = stdout.Capture(toolchain.PrintHelpMessage)
		expectedOutput = getExpectedPrintHelpMessage(toolchain)

		return actualOutput, expectedOutput
	}

	t.Run("should print the help message for the toolchain.", func(t *testing.T) {
		if actualOutput, expectedOutput := setup(); actualOutput != expectedOutput {
			t.Errorf("Expected output to be the following:\n %q\n\n But got the following instead:\n %q", expectedOutput, actualOutput)
		}
	})
}

func TestRunCommand(t *testing.T) {
	setup := func(commandName string) (actualOutput, expectedOutput string, expectedErrorMessage string, err error) {
		toolchain := NewToolchain(WithCommand(mc.NewMockCommand()))

		stdout := utils.NewPrintStorage()
		actualOutput, err = stdout.CaptureWithError(func() error {
			return toolchain.RunCommand(commandName, nil)
		})

		expectedErrorMessage = fmt.Sprintf("command '%v' is not valid", f.Invalid)
		expectedOutput = getExpectedPrintHelpMessage(toolchain)

		return actualOutput, expectedOutput, expectedErrorMessage, err
	}

	t.Run("Given a valid command name as an argument", func(t *testing.T) {
		t.Run("When its called with the 'help' command name", func(t *testing.T) {
			t.Parallel()
			t.Run("Then it should print the help message to the stdout.", func(t *testing.T) {
				if actualOutput, expectedOutput, _, _ := setup("help"); actualOutput != expectedOutput {
					t.Errorf("Expected output to be the following:\n %q\n\n But got the following instead:\n %q", expectedOutput, actualOutput)
				}
			})
		})

		t.Run("When its called with (another) valid command name", func(t *testing.T) {
			t.Parallel()
			t.Run("Then it should call the provided mock command executor to return a nil error value.", func(t *testing.T) {
				if _, _, _, err := setup("mock"); err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
			})
		})
	})

	t.Run("Given an invalid command name as an argument", func(t *testing.T) {
		t.Parallel()
		t.Run("When its called with a nil/empty string value in place of a command name", func(t *testing.T) {
			t.Run("Then it should do nothing by returning a nil error value", func(t *testing.T) {
				t.Parallel()
				if _, _, _, err := setup(""); err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
			})

			t.Run("And it should print nothing to the stdout.", func(t *testing.T) {
				t.Parallel()
				if actualOutput, _, _, _ := setup(""); actualOutput != "" {
					t.Errorf("Expected output to be the following:\n %q\n\n But got the following instead:\n %q", "", actualOutput)
				}
			})
		})

		t.Run("When its called with an invalid command name", func(t *testing.T) {
			t.Parallel()
			t.Run("Then it should return a non-nil error", func(t *testing.T) {
				t.Parallel()
				if _, _, _, err := setup(f.Invalid); err == nil {
					t.Error("Expected a non-nil error value.")
				}
			})

			t.Run("And it should return with the correct error message.", func(t *testing.T) {
				t.Parallel()
				if _, _, expectedErrorMessage, err := setup(f.Invalid); err.Error() != expectedErrorMessage {
					t.Errorf("Expected error to be: %v, but got: %v", expectedErrorMessage, err.Error())
				}
			})
		})
	})
}

func TestToolchain_SelectCommand(t *testing.T) {
	const (
		invalidCommand = "nonexistent"
		validCommand   = "mock1"
	)

	runSelectCommandTest := func(commandNameQuery string) (foundCommand *c.Command, commandExists bool, expectedCommand *mc.MockCommand) {
		expectedCommand = mc.NewMockCommand(mc.WithName(validCommand))
		toolchain := NewToolchain(
			WithCommand(expectedCommand),
			WithCommand(mc.NewMockCommand(mc.WithName("mock2"))),
		)
		foundCommand, commandExists = toolchain.SelectCommand(commandNameQuery)
		return foundCommand, commandExists, expectedCommand
	}

	t.Run("Given a VALID command name as it's argument", func(t *testing.T) {
		t.Parallel()
		t.Run("should select the correct command when queried by name.", func(t *testing.T) {
			t.Parallel()
			if foundCommand, _, expectedCommand := runSelectCommandTest(validCommand); (*foundCommand).GetName() != expectedCommand.GetName() {
				t.Errorf("Expected foundCommand to be command1, but got %v", foundCommand)
			}
		})

		t.Run("should return a true boolean value if the command exists.", func(t *testing.T) {
			t.Parallel()
			if _, commandExists, _ := runSelectCommandTest(validCommand); !commandExists {
				t.Errorf("Expected commandExists to be true, but got false")
			}
		})
	})

	t.Run("Given an INVALID command name as it's argument", func(t *testing.T) {
		t.Parallel()
		t.Run("should return a false boolean value if the command does not exist.", func(t *testing.T) {
			if _, commandExists, _ := runSelectCommandTest(invalidCommand); commandExists {
				t.Errorf("Expected commandExists to be false, but got true")
			}
		})
	})
}

func TestToolchain_addEmptyLine(t *testing.T) {
	runAddEmptyLineTest := func() (output string) {
		stdout := utils.NewPrintStorage()
		toolchain := &Toolchain{}
		output = stdout.Capture(toolchain.addEmptyLine)
		return output
	}

	t.Run("should add an empty line to the given string.", func(t *testing.T) {
		if output := runAddEmptyLineTest(); output != f.EmptyLine {
			t.Errorf("Expected output to be an empty line (\\n), but instead got %v", output)
		}
	})
}

func TestToolchain_printAllCommandHelps(t *testing.T) {
	runPrintAllCommandHelpsTest := func() (output string, expectedOutput string) {
		stdout := utils.NewPrintStorage()

		const command1_name, command1_description = "Command 1", "Command 1 help message."
		command1 := createHelpCommand(command1_name, command1_description)

		const command2_name, command2_description = "Command 2", "Command 2 help message."
		command2 := createHelpCommand(command2_name, command2_description)

		toolchain := NewToolchain(
			WithCommand(command1),
			WithCommand(command2),
		)

		expectedOutput = fmt.Sprintf("%s: %s\n%s: %s\n", command1_name, command1_description, command2_name, command2_description)
		output = stdout.Capture(toolchain.printAllCommandHelps)

		return output, expectedOutput
	}
	t.Run("should print the help message for all commands in the toolchain.", func(t *testing.T) {
		if output, expectedOutput := runPrintAllCommandHelpsTest(); output != expectedOutput {
			t.Errorf("Expected output to be the following:\n %q\n\n But got the following instead:\n %q", expectedOutput, output)
		}
	})

}

func createHelpCommand(name string, description string) *mc.MockCommand {
	return mc.NewMockCommand(
		mc.WithName(name),
		mc.WithDescription(description),
	)
}

func endWithNewLine(output string) string {
	return fmt.Sprintf("%s%s", output, f.EmptyLine)
}

func getExpectedPrintHelpMessage(toolchain *Toolchain) string {
	expectedOutput := endWithNewLine(toolchain.prints.Title)
	expectedOutput += f.EmptyLine
	expectedOutput += endWithNewLine(toolchain.prints.CommandsSubTitle)
	expectedOutput += endWithNewLine(toolchain.prints.HR)
	expectedOutput += f.EmptyLine

	for _, command := range *toolchain.commands {
		name, description := (*command).GetName(), (*command).GetDescription()
		expectedOutput += fmt.Sprintf("%s: %s\n", name, description)
	}

	return expectedOutput
}
