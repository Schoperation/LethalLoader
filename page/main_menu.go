package page

import (
	"fmt"
	"schoperation/lethalloader/domain/config"
	"schoperation/lethalloader/domain/option"
	"schoperation/lethalloader/domain/profile"
)

type mainConfigUpdater interface {
	Read() (config.MainConfig, error)
	Write(mainConfig config.MainConfig) error
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

func (page MainMenuPage) Show() (option.OptionsResults, error) {
	clear()

	mainConfig, err := page.mainConfigUpdater.Read()
	if err != nil {
		return option.OptionsResults{}, err
	}

	profiles, err := page.profileManager.GetAll()
	if err != nil {
		return option.OptionsResults{}, err
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
	fmt.Print("C ) Create New Profile\n")
	fmt.Print("En) Edit nth Profile\n")
	fmt.Print("Dn) Delete nth Profile\n")
	fmt.Print("Q ) Quit\n")
	fmt.Print("\n")

	options := option.NewOptions(option.OptionsArgs{
		Options: map[string]option.CmdName{
			"Sn": option.TaskSwitchProfile,
			"C":  option.TaskCreateProfile,
			"En": option.PageProfileViewer,
			"Dn": option.TaskDeleteProfile,
			"Q":  option.TaskQuit,
		},
	})

	return options.TakeInput(), nil
}
