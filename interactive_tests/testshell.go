package main

import "github.com/robwc/netsh"

func main() {
	sh := netsh.NewShell("netsh> ")
	sh.Start()
}
