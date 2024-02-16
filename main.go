package main

import (
	"fmt"
	"schoperation/lethalloader/adapter/file"
	"schoperation/lethalloader/command"
	"schoperation/lethalloader/command/task"
	translator_config "schoperation/lethalloader/translator/config"
)

func main() {
	mainConfigDao := file.NewMainConfigDao()

	mainConfigTranslator := translator_config.NewMainConfigTranslator(mainConfigDao)

	firstTimeSetupTask := task.NewFirstTimeSetupTask()

	mainMenuCmd := command.NewMainMenuCommand(mainConfigTranslator, firstTimeSetupTask)

	err := mainMenuCmd.Run()
	if err != nil {
		fmt.Printf("Error occurred: %v", err)
	}
}
