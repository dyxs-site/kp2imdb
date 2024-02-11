//go:build windows
// +build windows

package util

import (
	"os"
	"os/exec"
)

func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
