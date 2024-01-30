package main

import (
	"fmt"
	"schoperation/lethalloader/command"
)

func main() {
	startupCommand := command.NewStartupCommand()

	err := startupCommand.Run()
	if err != nil {
		fmt.Printf("Error occurred")
	}
}
