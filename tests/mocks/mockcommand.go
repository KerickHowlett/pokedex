package mocks

import "fmt"

type MockCommand struct {
	Name        string
	Description string
}

func (c *MockCommand) GetName() string {
	return c.Name
}

func (c *MockCommand) SetName(name string) {
	c.Name = name
}

func (c *MockCommand) GetDescription() string {
	return c.Description
}

func (c *MockCommand) SetDescription(description string) {
	c.Description = description
}

func (c *MockCommand) Execute() error {
	return nil
}

func (c *MockCommand) PrintHelp() {
	name := c.GetName()
	description := c.GetDescription()

	fmt.Printf("%s: %s\n", name, description)
}

func NewMockCommand() *MockCommand {
	return &MockCommand{
		Name:        "mock",
		Description: "This is a mocked command.",
	}
}
