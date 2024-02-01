package main

import (
	"fmt"
	"schoperation/lethalloader/adapter/file"
	"schoperation/lethalloader/command"
	translator_config "schoperation/lethalloader/translator/config"
)

func main() {
	mainConfigDao := file.NewMainConfigDao()

	mainConfigTranslator := translator_config.NewMainConfigTranslator(mainConfigDao)

	startupCommand := command.NewStartupCommand(mainConfigTranslator)

	err := startupCommand.Run()
	if err != nil {
		fmt.Printf("Error occurred: %v", err)
	}
}
