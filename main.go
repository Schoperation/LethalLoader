package main

/*
   LethalLoader
   Copyright (C) 2024 Schoperation

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
	gameFilesTranslator := translator_mod.NewGameFilesTranslator(gameFilesDao, mainConfigDao)

	firstTimeSetupTask := task.NewFirstTimeSetupTask(mainConfigTranslator, gameFilesDao, profileTranslator)
	newProfileTask := task.NewNewProfileTask(profileTranslator)
	deleteProfileTask := task.NewDeleteProfileTask(profileTranslator)
	searchTermTask := task.NewSearchTermTask()
	addModTask := task.NewAddModToProfileTask(listingTranslator, modTranslator, gameFilesTranslator, profileTranslator)
	removeModTask := task.NewRemoveModTask(profileTranslator, gameFilesTranslator)
	updateModsTask := task.NewUpdateModsTask(modTranslator, profileTranslator)
	switchProfileTask := task.NewSwitchProfileTask(mainConfigTranslator, profileTranslator)

	mainMenuPage := page.NewMainMenuPage(mainConfigTranslator, profileTranslator)
	profileViewerPage := page.NewProfileViewerPage()
	modSearchResultsPage := page.NewModSearchResultsPage(searchResultTranslator, modTranslator)
	checkForModUpdatesPage := page.NewCheckForModUpdatesPage(listingTranslator)
	aboutPage := page.NewAboutPage()

	pageViewer := viewer.NewPageViewer(
		mainMenuPage,
		profileViewerPage,
		modSearchResultsPage,
		checkForModUpdatesPage,
		aboutPage,
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
