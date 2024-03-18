package mocks

import "fmt"

type MockCommand struct {
	Name        string
	Description string
}

func (c *MockCommand) GetName() string {
	return "mock"
}

func (c *MockCommand) GetDescription() string {
	return "This is a mocked command."
}

func (c *MockCommand) Execute() error {
	return nil
}

func (c *MockCommand) PrintHelp() {
	name := c.GetName()
	description := c.GetDescription()

	fmt.Printf("%s: %s\n", name, description)
}
