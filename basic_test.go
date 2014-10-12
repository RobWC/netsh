package netsh

import "testing"

func BasicShellTest(t *testing.T) {
	sh := NewShell("netsh> ")
	sh.Start()
}
