package page

import (
	"fmt"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/viewer"
)

type ModSearchResultsPage struct {
}

func NewModSearchResultsPage() ModSearchResultsPage {
	return ModSearchResultsPage{}
}

func (page ModSearchResultsPage) Show(args any) (viewer.OptionsResult, error) {
	clear()

	input, ok := args.(input.ModSearchResultsPageInput)
	if !ok {
		return viewer.OptionsResult{}, fmt.Errorf("could not parse input")
	}

	fmt.Printf("Adding Mod to %s\n", input.Profile.Name())
	fmt.Print("---------------------------------------\n\n")

	fmt.Printf("Search Results for %s:\n", input.Term)
}
