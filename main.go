package main

import (
	"fmt"
	"schoperation/lethalloader/adapter/file"
	"schoperation/lethalloader/page"
	"schoperation/lethalloader/task"
	translator_config "schoperation/lethalloader/translator/config"
	translator_profile "schoperation/lethalloader/translator/profile"
)

func main() {

	steamChecker := file.NewSteamChecker()
	mainConfigDao := file.NewMainConfigDao()
	profileDao := file.NewProfileDao()
	modListDao := file.NewModListDao()

	mainConfigTranslator := translator_config.NewMainConfigTranslator(mainConfigDao)
	profileTranslator := translator_profile.NewProfileTranslator(profileDao, modListDao)

	firstTimeSetupTask := task.NewFirstTimeSetupTask(mainConfigTranslator, steamChecker, profileTranslator)
	err := firstTimeSetupTask.Do()
	if err != nil {
		fmt.Printf("Failed to perform first time setup: %v\n", err)
		return
	}

	mainMenuPage := page.NewMainMenuPage(mainConfigTranslator, profileTranslator)

	pageViewer := NewPageViewer(
		mainMenuPage,
	)

	err = pageViewer.Run()
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	}
}
