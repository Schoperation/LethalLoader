package main

import (
	"fmt"
	"schoperation/lethalloader/adapter/file"
	"schoperation/lethalloader/adapter/rest"
	"schoperation/lethalloader/page"
	"schoperation/lethalloader/task"
	translator_config "schoperation/lethalloader/translator/config"
	translator_mod "schoperation/lethalloader/translator/mod"
	translator_profile "schoperation/lethalloader/translator/profile"
)

func main() {
	steamChecker := file.NewSteamChecker()
	mainConfigDao := file.NewMainConfigDao()
	profileDao := file.NewProfileDao()
	modListDao := file.NewModListDao()

	modDownloader := rest.NewModDownloader()

	mainConfigTranslator := translator_config.NewMainConfigTranslator(mainConfigDao)
	profileTranslator := translator_profile.NewProfileTranslator(profileDao, modListDao)
	modTranslator := translator_mod.NewModTranslator(modDownloader, modListDao)

	firstTimeSetupTask := task.NewFirstTimeSetupTask(mainConfigTranslator, steamChecker, profileTranslator, modTranslator)

	mainMenuPage := page.NewMainMenuPage(mainConfigTranslator, profileTranslator)

	pageViewer := NewPageViewer(
		mainMenuPage,
		firstTimeSetupTask,
	)

	err := pageViewer.Run()
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	}
}
