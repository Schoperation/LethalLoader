package page

import (
	"fmt"
	"schoperation/lethalloader/domain/option"
	"schoperation/lethalloader/domain/profile"
)

type ProfileViewerPage struct {
}

func NewProfileViewerPage() ProfileViewerPage {
	return ProfileViewerPage{}
}

func (page ProfileViewerPage) Show(args ...any) (option.OptionsResults, error) {
	clear()

	profile, ok := args[0].(profile.Profile)
	if !ok {
		return option.OptionsResults{}, fmt.Errorf("could not cast profile")
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

	options := option.NewOptions(option.NewOptionsArgs{
		Pages: map[string]option.PageName{
			"Q": option.PageMainMenu,
		},
		Tasks: map[string]option.TaskName{},
	})

	return options.TakeInput(), nil
}
