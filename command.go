package netsh

type Command struct {
	Name          string
	Description   string
	HelpText      string
	Limits        interface{}
	ChildCommands []Command
}
