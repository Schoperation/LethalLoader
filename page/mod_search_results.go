package page

import (
	"fmt"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type modSearcher interface {
	Search(term string) ([]mod.SearchResult, error)
}

type ModSearchResultsPage struct {
	modSearcher modSearcher
}

func NewModSearchResultsPage(
	modSearcher modSearcher,
) ModSearchResultsPage {
	return ModSearchResultsPage{
		modSearcher: modSearcher,
	}
}

func (page ModSearchResultsPage) Show(args any) (viewer.OptionsResult, error) {
	clear()

	input, ok := args.(input.ModSearchResultsPageInput)
	if !ok {
		return viewer.OptionsResult{}, fmt.Errorf("could not parse input")
	}

	searchResults, err := page.modSearcher.Search(input.Term)
	if err != nil {
		return viewer.OptionsResult{}, err
	}

	fmt.Printf("Adding Mod to %s\n", input.Profile.Name())
	fmt.Print("---------------------------------------\n\n")

	if len(searchResults) == 0 {
		fmt.Printf("No results found for %s.\n", input.Term)

		searchTermOp := viewer.NewOption(viewer.OptionDto{
			Letter: 'A',
			Task:   viewer.TaskSearchTerm,
		}, []profile.Profile{input.Profile})

		return viewer.NewOptionsResult(searchTermOp, -1)
	}

	fmt.Printf("Search Results for %s:\n", input.Term)
	for i, result := range searchResults {
		fmt.Printf("\t%d ~ %s ~ %s ~ %s\n", i+1, result.Name(), result.Author(), result.Description())
	}

	fmt.Print("\n")
	fmt.Print("What to Do?\n")
	fmt.Print("-----------\n")
	fmt.Print("An) Add nth Mod\n")
	fmt.Print("S) Search New Term\n")
	fmt.Print("Q) Cancel\n")
	fmt.Print("\n")

}

func (page ModSearchResultsPage) options(pfToView profile.Profile) viewer.Options {
	cancel := viewer.NewOption(viewer.OptionDto{
		Letter: 'Q',
		Page:   viewer.PageProfileViewer,
	}, []profile.Profile{pfToView})

	return viewer.NewOptions(
		[]viewer.Option{
			cancel,
		},
	)
}
