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
	Tab     = []byte{99}
	NewLine = []byte{10}
	Space   = []byte{32}
)

type Shell struct {
	prompt      string
	currentLine []byte
}

func NewShell(prompt string) *Shell {
	return &Shell{prompt: prompt, currentLine: make([]byte, 0)}
}

func (sh *Shell) SetPrompt(prompt string) {
	sh.prompt = prompt
}

func (sh *Shell) processCommands() {

}

func (sh *Shell) processChar(b []byte) {
	sh.currentLine = append(currentLine, b[0])
	if bytes.Compare(b, Tab) == 0 {
		//tab
		fmt.Println("TAB")
	} else if bytes.Compare(b, NewLine) == 0 {
		//newline
		fmt.Println("N")
		fmt.Println(currentLine)
		fmt.Println("I got the byte", b, "("+string(b)+")")
		fmt.Printf("%s", sh.prompt)
		currentLine = make([]byte, 0)
	} else if bytes.Compare(b, Space) == 0 {
		//space

	}
}

func (sh *Shell) Start() {

}

func (sh *Shell) Exit() {

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

func main() {
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
