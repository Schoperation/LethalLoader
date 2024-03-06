package page

import (
	"fmt"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type ProfileViewerPage struct {
}

func NewProfileViewerPage() ProfileViewerPage {
	return ProfileViewerPage{}
}

func (page ProfileViewerPage) Show(args ...any) (viewer.OptionsResult, error) {
	clear()

	profile, ok := args[0].(profile.Profile)
	if !ok {
		return viewer.OptionsResult{}, fmt.Errorf("could not cast profile")
	}

	fmt.Printf("Profile %s\n", profile.Name())
	fmt.Print("---------------------------------------\n\n")

	for i, mod := range profile.Mods() {
		fmt.Printf("\t%d ~ %s ~ v%s ~ by %s ~ %s\n", i+1, mod.Name(), mod.Version(), mod.Author(), mod.Description())
	}

	fmt.Print("\n")
	fmt.Print("What to Do?\n")
	fmt.Print("-----------\n")
	fmt.Print("U) Check for Mod Updates\n")
	fmt.Print("A) Add Mod\n")
	fmt.Print("R) Remove Mod\n")
	fmt.Print("Q) Quit to Main Menu\n")
	fmt.Print("\n")

	options := page.options()

	return options.TakeInput(), nil
}

func (page ProfileViewerPage) options() viewer.Options {
	quit := viewer.NewOption(viewer.OptionDto{
		Letter: 'Q',
		Task:   viewer.TaskQuit,
	}, []string{})

	return viewer.NewOptions(
		[]viewer.Option{
			quit,
		},
	)
}
