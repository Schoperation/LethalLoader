package page

import (
	"fmt"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type ProfileViewerPage struct {
}

func NewProfileViewerPage() ProfileViewerPage {
	return ProfileViewerPage{}
}

func (page ProfileViewerPage) Show(args any) (viewer.OptionsResult, error) {
	clear()

	pfToView, ok := args.(profile.Profile)
	if !ok {
		return viewer.OptionsResult{}, fmt.Errorf("could not cast profile")
	}

	fmt.Printf("Profile %s\n", pfToView.Name())
	fmt.Print("---------------------------------------\n\n")

	for i, mod := range pfToView.Mods() {
		fmt.Printf("\t%02d ~ %s ~ v%s ~ by %s ~ %s\n", i+1, mod.Name(), mod.Version(), mod.Author(), mod.Description())
	}

	fmt.Print("\n")
	fmt.Print("What to Do?\n")
	fmt.Print("-----------\n")
	fmt.Print("U) Check for Mod Updates\n")
	fmt.Print("A) Add Mod\n")
	fmt.Print("R) Remove Mod\n")
	fmt.Print("Q) Back to Main Menu\n")
	fmt.Print("\n")

	options := page.options(pfToView)
	return options.TakeInput(), nil
}

func (page ProfileViewerPage) options(pfToView profile.Profile) viewer.Options {
	checkForUpdates := viewer.NewOption(viewer.OptionDto{
		Letter: 'U',
		Page:   viewer.PageCheckForModUpdates,
	}, []profile.Profile{pfToView})

	addMod := viewer.NewOption(viewer.OptionDto{
		Letter: 'A',
		Task:   viewer.TaskSearchTerm,
	}, []input.SearchTermTaskInput{
		{
			Profile:         pfToView,
			SkipCacheSearch: false,
		}})

	removeModArgs := make([]input.RemoveModTaskInput, len(pfToView.Mods()))
	for i, mod := range pfToView.Mods() {
		removeModArgs[i] = input.RemoveModTaskInput{
			Profile: pfToView,
			Mod:     mod,
		}
	}

	removeMod := viewer.NewOption(viewer.OptionDto{
		Letter:   'R',
		Task:     viewer.TaskRemoveMod,
		TakesNum: true,
	}, removeModArgs)

	quit := viewer.NewOption(viewer.OptionDto{
		Letter: 'Q',
		Page:   viewer.PageMainMenu,
	}, []string{})

	return viewer.NewOptions(
		[]viewer.Option{
			checkForUpdates,
			addMod,
			removeMod,
			quit,
		},
	)
}
