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

	const characterLimit = 20
	longestNameChars := 0
	longestAuthorChars := 0
	for _, mod := range pfToView.Mods() {
		if len(mod.Name()) > longestNameChars {
			longestNameChars = len(mod.Name())
		}

		if len(mod.Author()) > longestAuthorChars {
			longestAuthorChars = len(mod.Author())
		}
	}

	if longestNameChars > characterLimit {
		longestNameChars = characterLimit
	}

	if longestAuthorChars > characterLimit {
		longestAuthorChars = characterLimit
	}

	fmt.Printf("Profile %s ~ %d Mods\n", pfToView.Name(), pfToView.NumberOfMods())
	fmt.Print("---------------------------------------\n\n")

	for i, mod := range pfToView.Mods() {
		fmt.Printf("\t%02d ~ %s\n", i+1, mod.PrettyPrint(longestNameChars, longestAuthorChars))
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
