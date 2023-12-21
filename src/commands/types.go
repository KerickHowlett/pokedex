package commands

type CLICommand struct {
	Name        string
	Description string
	Callback    func() error
}
