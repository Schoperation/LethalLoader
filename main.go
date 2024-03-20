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
	"schoperation/lethalloader/viewer"
)

func main() {
	mainConfigDao := file.NewMainConfigDao()
	profileDao := file.NewProfileDao()
	modListDao := file.NewModListDao()
	gameFilesDao := file.NewGameFilesDao()
	unzipper := file.NewFileUnzipper()

	modDownloader := rest.NewModDownloader()
	thunderstoreClient := rest.NewThunderstoreClient()

	mainConfigTranslator := translator_config.NewMainConfigTranslator(mainConfigDao)
	profileTranslator := translator_profile.NewProfileTranslator(profileDao, modListDao, gameFilesDao)
	modTranslator := translator_mod.NewModTranslator(modDownloader, unzipper, modListDao)
	listingTranslator := translator_mod.NewListingTranslator(thunderstoreClient)
	searchResultTranslator := translator_mod.NewSearchResultTranslator(thunderstoreClient)

	firstTimeSetupTask := task.NewFirstTimeSetupTask(mainConfigTranslator, gameFilesDao, profileTranslator)
	newProfileTask := task.NewNewProfileTask(profileTranslator)
	deleteProfileTask := task.NewDeleteProfileTask(profileTranslator)
	searchTermTask := task.NewSearchTermTask()
	addModTask := task.NewAddModToProfileTask(listingTranslator, modTranslator, profileTranslator, mainConfigTranslator)
	removeModTask := task.NewRemoveModTask(profileTranslator, mainConfigTranslator)
	updateModsTask := task.NewUpdateModsTask(modTranslator, profileTranslator)
	switchProfileTask := task.NewSwitchProfileTask(mainConfigTranslator, profileTranslator)

	mainMenuPage := page.NewMainMenuPage(mainConfigTranslator, profileTranslator)
	profileViewerPage := page.NewProfileViewerPage()
	modSearchResultsPage := page.NewModSearchResultsPage(searchResultTranslator, modTranslator)
	checkForModUpdatesPage := page.NewCheckForModUpdatesPage(listingTranslator)

	pageViewer := viewer.NewPageViewer(
		mainMenuPage,
		profileViewerPage,
		modSearchResultsPage,
		checkForModUpdatesPage,
		firstTimeSetupTask,
		newProfileTask,
		deleteProfileTask,
		searchTermTask,
		addModTask,
		removeModTask,
		updateModsTask,
		switchProfileTask,
	)

	err := pageViewer.Run()
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	}
}
