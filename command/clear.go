package command

import (
	"os"
	"os/exec"
	"runtime"
)

func clear() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls")
	}

	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
