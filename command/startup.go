package command

import (
	"fmt"
	"schoperation/lethalloader/domain/config"
)

type mainConfigReader interface {
	Read() (config.MainConfig, error)
}

type StartupCommand struct {
	mainConfigReader mainConfigReader
}

func NewStartupCommand(
	mainConfigReader mainConfigReader,
) StartupCommand {
	return StartupCommand{
		mainConfigReader: mainConfigReader,
	}
}

func (cmd StartupCommand) Run() error {
	fmt.Print("LethalLoader v0.0.1 ALPHA (expect bugs)\n")
	fmt.Print("---------------------------------------\n\n")

	mainConfig, err := cmd.mainConfigReader.Read()
	if err != nil {
		return err
	}

	if mainConfig.GameFilePath() == "" {
		fmt.Print("It appears this is your first time. Setting up...\n")
	}

	return nil
}
