package command

import (
	"fmt"
	"schoperation/lethalloader/domain/config"
)

type mainConfigReader interface {
	Read() (config.MainConfig, error)
}

type firstTimeSetupTask interface {
	Do() error
}

type MainMenuCommand struct {
	mainConfigReader   mainConfigReader
	firstTimeSetupTask firstTimeSetupTask
}

func NewMainMenuCommand(
	mainConfigReader mainConfigReader,
	firstTimeSetupTask firstTimeSetupTask,
) MainMenuCommand {
	return MainMenuCommand{
		mainConfigReader:   mainConfigReader,
		firstTimeSetupTask: firstTimeSetupTask,
	}
}

func (cmd MainMenuCommand) Run() error {
	clear()

	fmt.Print("LethalLoader v0.0.1 ALPHA (expect bugs)\n")
	fmt.Print("---------------------------------------\n\n")

	mainConfig, err := cmd.mainConfigReader.Read()
	if err != nil {
		return err
	}

	if mainConfig.GameFilePath() == "" {
		fmt.Print("It appears this is your first time. Setting up...\n")

		err := cmd.firstTimeSetupTask.Do()
		if err != nil {
			return err
		}
	}

	return nil
}
