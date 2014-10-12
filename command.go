package netsh

type CommandActions interface {
	Run()
}

type Command struct {
	Name        string
	Description string
	HelpText    string
	Limits      interface{}
}
