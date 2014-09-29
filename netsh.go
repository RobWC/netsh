package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

var (
	Tab     = []byte{99} //Tab byre
	NewLine = []byte{10} //New Line byte
	Space   = []byte{32} //Space byte
)

type Command struct {
	Name        string
	Description string
	HelpText    string
	Limits      interface{}
}

//Shell
type Shell struct {
	prompt      string
	currentLine []byte
}

//NewShell
func NewShell(prompt string) *Shell {
	return &Shell{prompt: prompt, currentLine: make([]byte, 0)}
}

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
		sh.currentLine = make([]byte, 0)
	}
}

func (sh *Shell) processChar(b []byte) {
	if bytes.Compare(b, Tab) == 0 {
		//tab
		fmt.Println("TAB")
		fmt.Print(string(b))
	} else if bytes.Compare(b, NewLine) == 0 {
		fmt.Print(string(b))
		//fmt.Println("I got the byte", b, "("+string(b)+")")
		sh.processCommands()
		sh.outputPrompt()
	} else if bytes.Compare(b, Space) == 0 {
		//space
		sh.currentLine = append(sh.currentLine, b[0])
		fmt.Print(string(b))
	} else {
		sh.currentLine = append(sh.currentLine, b[0])
		fmt.Print(string(b))
	}
}

func (sh *Shell) outputPrompt() {
	fmt.Printf("%s", sh.prompt)
}

func (sh *Shell) Start() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			// sig is a ^C, handle it
			fmt.Println(sig)
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()
			os.Exit(0)
		}
	}()
	sh.outputPrompt()
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		sh.processChar(b)

	}
}

func (sh *Shell) Exit() {
	//exit or shutdown
}

func main() {
	sh := NewShell("netsh> ")
	sh.Start()
}

var prompt string
var err error
var currentLine []byte

func main1() {
	prompt = "netsh> "
	fmt.Printf("%s", prompt)
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strings.TrimSuffix(input, "\n"), len(input))
}

func processChar(b []byte) {
	prompt = "netsh> "
	currentLine = append(currentLine, b[0])
	if bytes.Compare(b, []byte{9}) == 0 {
		//tab
		fmt.Println("TAB")
	} else if bytes.Compare(b, []byte{10}) == 0 {
		//newline
		fmt.Println("N")
		fmt.Println(currentLine)
		fmt.Println("I got the byte", b, "("+string(b)+")")
		fmt.Printf("%s", prompt)
		currentLine = make([]byte, 0)
	} else if bytes.Compare(b, []byte{32}) == 0 {
		//space

	}
}

func main2() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			// sig is a ^C, handle it
			fmt.Println(sig)
			exec.Command("stty", "-F", "/dev/tty", "echo")
			os.Exit(0)
		}
	}()
	prompt = "netsh> "
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo")

	var b []byte = make([]byte, 1)
	fmt.Printf("%s", prompt)
	for {
		os.Stdin.Read(b)
		processChar(b)

	}
}
