package page

import (
	"fmt"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type profileImporter interface {
	GetAll() ([]profile.ExportedProfile, error)
}

type ImportProfilePage struct {
	profileImporter profileImporter
}

func NewImportProfilePage(
	profileImporter profileImporter,
) ImportProfilePage {
	return ImportProfilePage{
		profileImporter: profileImporter,
	}
}

func (page ImportProfilePage) Show(args any) (viewer.OptionsResult, error) {
	clear()

	fmt.Print("Import Profile\n")
	fmt.Print("---------------------------------------\n\n")

	profiles, err := page.profileImporter.GetAll()
	if err != nil {
		return viewer.OptionsResult{}, err
	}

	for i, pf := range profiles {
		fmt.Printf("\t%02d ~ %s ~ %d mods\n", i+1, pf.Name(), len(pf.ModSlugs()))
	}

	if len(profiles) == 0 {
		fmt.Printf("\tNo profiles found. Put them in the profiles folder, wherever you've put your LethalLoader files.\n\n")
	}

	fmt.Print("\n")
	fmt.Print("What to Do?\n")
	fmt.Print("-----------\n")
	fmt.Print("In) Import nth Profile\n")
	fmt.Print("R ) Refresh Files\n")
	fmt.Print("Q ) Back to Main Menu\n")
	fmt.Print("\n")

	options := page.options(profiles)
	return options.TakeInput(), nil
}

func (page ImportProfilePage) options(profiles []profile.ExportedProfile) viewer.Options {
	importProfile := viewer.NewOption(viewer.OptionDto{
		Letter:   'I',
		Task:     viewer.TaskImportProfile,
		TakesNum: true,
	}, profiles)

	refreshFiles := viewer.NewOption(viewer.OptionDto{
		Letter: 'R',
		Page:   viewer.PageImportProfile,
	}, []string{})

	quit := viewer.NewOption(viewer.OptionDto{
		Letter: 'Q',
		Page:   viewer.PageMainMenu,
	}, []string{})

	return viewer.NewOptions(
		[]viewer.Option{
			importProfile,
			refreshFiles,
			quit,
		},
	)
}
