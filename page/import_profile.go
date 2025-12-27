package page

import (
	"fmt"
	"schoperation/lethalloader/domain/viewer"
)

type ImportProfilePage struct {
}

func NewImportProfilePage() ImportProfilePage {
	return ImportProfilePage{}
}

func (page ImportProfilePage) Show(args any) (viewer.OptionsResult, error) {
	clear()

	fmt.Print("\n")
	fmt.Print("What to Do?\n")
	fmt.Print("-----------\n")
	fmt.Print("In) Import nth Profile\n")
	fmt.Print("R ) Refresh Files\n")
	fmt.Print("Q ) Back to Main Menu\n")
	fmt.Print("\n")

	options := page.options()
	return options.TakeInput(), nil
}

func (page ImportProfilePage) options() viewer.Options {
	importProfile := viewer.NewOption(viewer.OptionDto{
		Letter:   'I',
		Task:     viewer.TaskAddMod,
		TakesNum: true,
	}, []string{})

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
