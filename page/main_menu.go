package page

import (
	"fmt"
	"schoperation/lethalloader/domain/config"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type mainConfigUpdater interface {
	Get() (config.MainConfig, error)
	Save(mainConfig config.MainConfig) error
}

type profileManager interface {
	GetAll() ([]profile.Profile, error)
}

type MainMenuPage struct {
	mainConfigUpdater mainConfigUpdater
	profileManager    profileManager
}

func NewMainMenuPage(
	mainConfigUpdater mainConfigUpdater,
	profileManager profileManager,
) MainMenuPage {
	return MainMenuPage{
		mainConfigUpdater: mainConfigUpdater,
		profileManager:    profileManager,
	}
}

func (page MainMenuPage) Show(args any) (viewer.OptionsResult, error) {
	clear()

	mainConfig, err := page.mainConfigUpdater.Get()
	if err != nil {
		return viewer.OptionsResult{}, err
	}

	profiles, err := page.profileManager.GetAll()
	if err != nil {
		return viewer.OptionsResult{}, err
	}

	fmt.Print("LethalLoader v0.0.1 ALPHA (expect bugs)\n")
	fmt.Print("---------------------------------------\n\n")

	fmt.Print("Profiles:\n")
	for i, pf := range profiles {
		fmt.Printf("\t%d ~ %s", i+1, pf.Name())

		if pf.Name() == mainConfig.SelectedProfile() {
			fmt.Print(" *SELECTED*")
		}

		fmt.Print("\n")
	}

	fmt.Print("\n")
	fmt.Print("What to Do?\n")
	fmt.Print("-----------\n")
	fmt.Print("Sn) Switch to nth Profile\n")
	fmt.Print("N ) New Profile\n")
	fmt.Print("En) Edit nth Profile\n")
	fmt.Print("Dn) Delete nth Profile\n")
	fmt.Print("Q ) Quit\n")
	fmt.Print("\n")

	options := page.options(profiles)

	return options.TakeInput(), nil
}

func (page MainMenuPage) options(profiles []profile.Profile) viewer.Options {
	switchProfile := viewer.NewOption(viewer.OptionDto{
		Letter:   'S',
		Task:     "task",
		TakesNum: true,
	}, profiles)

	newProfile := viewer.NewOption(viewer.OptionDto{
		Letter: 'N',
		Task:   viewer.TaskNewProfile,
	}, []string{})

	editProfile := viewer.NewOption(viewer.OptionDto{
		Letter:   'E',
		Page:     viewer.PageProfileViewer,
		TakesNum: true,
	}, profiles)

	deleteProfile := viewer.NewOption(viewer.OptionDto{
		Letter:   'D',
		Task:     viewer.TaskDeleteProfile,
		TakesNum: true,
	}, profiles)

	quit := viewer.NewOption(viewer.OptionDto{
		Letter: 'Q',
		Task:   viewer.TaskQuit,
	}, []string{})

	return viewer.NewOptions(
		[]viewer.Option{
			switchProfile,
			newProfile,
			editProfile,
			deleteProfile,
			quit,
		},
	)
}
