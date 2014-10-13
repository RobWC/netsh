package netsh

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"sort"
	"strings"
)

//Shell A shell that is used to process user input and commands
//Base structure for the shell
type Shell struct {
	prompt       string
	currentLine  []byte
	commands     map[string]Command
	commandOrder []string
}

//NewShell Create an initialized shell struct
func NewShell(prompt string) *Shell {
	var newCommands map[string]Command
	newCommands = make(map[string]Command)
	var order []string

	newCommands["foo"] = Command{Name: "foo", Description: "foo1", HelpText: "Eat the foo at the bar"}
	newCommands["bar"] = Command{Name: "bar", Description: "bar1", HelpText: "Eat the bar at the show"}
	newCommands["foof"] = Command{Name: "foof", Description: "foof1", HelpText: "Eat the foof"}

	for key := range newCommands {
		order = append(order, key)
	}
	sort.Strings(order)
	return &Shell{prompt: prompt, currentLine: make([]byte, 0), commandOrder: order, commands: newCommands}
}

//SetPrompt set the prompt of the shell
func (sh *Shell) SetPrompt(prompt string) {
	sh.prompt = prompt
}

//processCommands
// Takes in current commands and parses it.
// If the command is valid it is executed.
// If the command is invalid the command help is returned
// If the command is incomple, complete it
func (sh *Shell) processCommands(clear bool) {
	commands := strings.Split(string(sh.currentLine), " ")
	if len(commands) > 0 {
		re := regexp.MustCompile(regexp.QuoteMeta(commands[0]) + ".*")
		for key := range sh.commandOrder {
			if re.FindString(sh.commands[sh.commandOrder[key]].Name) != "" {
				fmt.Println(sh.commands[sh.commandOrder[key]].Name, "-", sh.commands[sh.commandOrder[key]].HelpText)
			}
		}
	}
	if clear {
		sh.currentLine = make([]byte, 0)
	}

}

func (sh *Shell) nextCommand(command []rune) {

}

func (sh *Shell) processChar(b byte) {
	//fmt.Println("Byte", b)
	if b == Tab {
		//tab
		//fmt.Print(string(b))
		//sh.processCommands(false)
		//sh.outputPrompt()
		//sh.outputCurrentline()
	} else if b == NewLine {
		//sh.processCommands(true)
		//sh.outputPrompt()
		//sh.outputCurrentline()
	} else if b == Space {
		//space
		sh.currentLine = append(sh.currentLine, b)
		//fmt.Print(string(b))
	} else {
		sh.currentLine = append(sh.currentLine, b)
		fmt.Print(string(b))
	}
}

//outputPrompt print a new prompt to stdoutF
func (sh *Shell) outputPrompt() {
	fmt.Printf("%s", sh.prompt)
}

func (sh *Shell) outputCurrentline() {
	fmt.Println("FOOD", string(sh.currentLine[:len(sh.currentLine)]))
	fmt.Printf("%s%s", sh.prompt, string(sh.currentLine[:]))
}

//Start start the shell in interactive mode
func (sh *Shell) Start() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	reader := bufio.NewReader(os.Stdin)

	//handle Ctrl-C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			// sig is a ^C, handle it
			fmt.Println(sig)
			sh.Exit()
		}
	}()

	sh.outputPrompt()
	for {
		r, _ := reader.ReadByte()
		sh.processChar(r)

	}
}

//Exit Exit the shell
func (sh *Shell) Exit() {
	//exit or shutdown
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	os.Exit(0)
}
