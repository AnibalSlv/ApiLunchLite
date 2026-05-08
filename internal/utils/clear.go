package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls") // Command Windows
	} else {
		cmd = exec.Command("clear") // Command Linux/macOS
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
