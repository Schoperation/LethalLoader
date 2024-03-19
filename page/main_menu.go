package page

import (
	"fmt"
	"schoperation/lethalloader/domain/config"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type selectedProfileGetter interface {
	Get() (config.MainConfig, error)
}

type allProfilesGetter interface {
	GetAll() ([]profile.Profile, error)
}

type MainMenuPage struct {
	selectedProfileGetter selectedProfileGetter
	allProfilesGetter     allProfilesGetter
}

func NewMainMenuPage(
	selectedProfileGetter selectedProfileGetter,
	allProfilesGetter allProfilesGetter,
) MainMenuPage {
	return MainMenuPage{
		selectedProfileGetter: selectedProfileGetter,
		allProfilesGetter:     allProfilesGetter,
	}
}

func (page MainMenuPage) Show(args any) (viewer.OptionsResult, error) {
	clear()

	mainConfig, err := page.selectedProfileGetter.Get()
	if err != nil {
		return viewer.OptionsResult{}, err
	}

	profiles, err := page.allProfilesGetter.GetAll()
	if err != nil {
		return viewer.OptionsResult{}, err
	}

	currentProfile := profile.Profile{}

	fmt.Print("LethalLoader v0.0.1 ALPHA (expect bugs)\n")
	fmt.Print("---------------------------------------\n\n")

	fmt.Print("Profiles:\n")
	for i, pf := range profiles {
		fmt.Printf("\t%02d ~ %s", i+1, pf.Name())

		if pf.Name() == mainConfig.SelectedProfile() {
			fmt.Print(" *SELECTED*")
			currentProfile = pf
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

	options := page.options(currentProfile, profiles)
	return options.TakeInput(), nil
}

func (page MainMenuPage) options(currentProfile profile.Profile, profiles []profile.Profile) viewer.Options {
	switchProfileInputs := make([]input.SwitchProfileTaskInput, len(profiles))
	for i, pf := range profiles {
		switchProfileInputs[i] = input.SwitchProfileTaskInput{
			OldProfile: currentProfile,
			NewProfile: pf,
		}
	}

	switchProfile := viewer.NewOption(viewer.OptionDto{
		Letter:   'S',
		Task:     viewer.TaskSwitchProfile,
		TakesNum: true,
	}, switchProfileInputs)

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
