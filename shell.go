package netsh

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

//Shell A shell that is used to process user input and commands
//Base structure for the shell
type Shell struct {
	prompt      string
	currentLine []rune
	commands    []Command
}

//NewShell Create an initialized shell struct
func NewShell(prompt string) *Shell {
	var newCommands []Command
	newCommands[0] = Command{Name: "foo", Description: "bar", HelpText: "Eat the foo at the bar"}
	return &Shell{prompt: prompt, currentLine: make([]rune, 0), commands: newCommands}
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
func (sh *Shell) processCommands() {
	fmt.Println(string(sh.currentLine))
	commands := strings.Split(string(sh.currentLine), " ")
	for item := range commands {
		fmt.Println("Command", commands[item])
		sh.currentLine = make([]rune, 0)
	}
}

func (sh *Shell) nextCommand(command []rune) {

}

func (sh *Shell) processChar(b rune) {
	if b == Tab {
		//tab
		//fmt.Print(string(b))
	} else if b == NewLine {
		fmt.Print(string(b))
		sh.processCommands()
		sh.outputPrompt()
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
	var b []byte
	for {
		os.Stdin.Read(b)
		r, _, _ := reader.ReadRune()
		sh.processChar(r)

	}
}

//Exit Exit the shell
func (sh *Shell) Exit() {
	//exit or shutdown
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	os.Exit(0)
}
