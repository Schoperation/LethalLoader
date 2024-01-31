package command

import (
	"os/exec"
	"runtime"
)

func clear() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls")
	}

	_ = cmd.Run()
}
