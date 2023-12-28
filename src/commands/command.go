package commands

type Command struct {
	Name        string
	Description string
	Execute     func() error
}
