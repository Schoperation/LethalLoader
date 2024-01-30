package command

import "fmt"

type StartupCommand struct {
}

func NewStartupCommand() StartupCommand {
	return StartupCommand{}
}

func (cmd StartupCommand) Run() error {
	fmt.Print("LethalLoader v0.0.1 ALPHA (expect bugs)\n")
	fmt.Print("---------------------------------------\n\n")

	return nil
}
