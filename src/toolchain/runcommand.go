package toolchain

import "fmt"

func (t *Toolchain) RunCommand(selection string) error {
	if selection == "" {
		return nil
	}

	if selection == "help" {
		t.PrintHelpMessage()
		return nil
	}

	if selectedCommand, commandFound := t.SelectCommand(selection); commandFound {
		return (*selectedCommand).Execute()
	}

	return fmt.Errorf("[ERROR] Command '%s' is not valid", selection)
}
