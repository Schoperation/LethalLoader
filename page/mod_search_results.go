package page

import (
	"fmt"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
	"strings"
)

type modSearcher interface {
	Search(term string) ([]mod.SearchResult, error)
}

type modCacheSearcher interface {
	GetAllBySearchTerm(term string) ([]mod.Mod, error)
}

type ModSearchResultsPage struct {
	modSearcher      modSearcher
	modCacheSearcher modCacheSearcher
}

func NewModSearchResultsPage(
	modSearcher modSearcher,
	modCacheSearcher modCacheSearcher,
) ModSearchResultsPage {
	return ModSearchResultsPage{
		modSearcher:      modSearcher,
		modCacheSearcher: modCacheSearcher,
	}
}

func (page ModSearchResultsPage) Show(args any) (viewer.OptionsResult, error) {
	clear()

	pageInput, ok := args.(input.ModSearchResultsPageInput)
	if !ok {
		return viewer.OptionsResult{}, fmt.Errorf("could not parse input")
	}

	fmt.Print("Searching...\n")

	cachedResults, err := page.modCacheSearcher.GetAllBySearchTerm(pageInput.Term)
	if err != nil {
		return viewer.OptionsResult{}, err
	}

	if len(cachedResults) > 0 && !pageInput.SkipCacheSearch {
		options := page.showCachedResults(pageInput, cachedResults)
		return options.TakeInput(), nil
	}

	searchResults, err := page.modSearcher.Search(pageInput.Term)
	if err != nil {
		return viewer.OptionsResult{}, err
	}

	if len(searchResults) == 0 {

		fmt.Printf("No Thunderstore results found for %s.\n", pageInput.Term)
		fmt.Printf("Try %s using spaces? Thunderstore's search is pretty bad...\n", page.termHasSpaces(pageInput.Term))

		searchTermOp := viewer.NewOption(viewer.OptionDto{
			Letter: 'A',
			Task:   viewer.TaskSearchTerm,
		}, []input.SearchTermTaskInput{{Profile: pageInput.Profile}})

		return viewer.NewOptionsResult(searchTermOp, -1)
	}

	options := page.showSearchResults(pageInput, searchResults)
	return options.TakeInput(), nil
}

func (page ModSearchResultsPage) showCachedResults(pageInput input.ModSearchResultsPageInput, mods []mod.Mod) viewer.Options {
	clear()

	fmt.Printf("Adding Mod to %s\n", pageInput.Profile.Name())
	fmt.Print("---------------------------------------\n\n")

	fmt.Printf("Cached Search Results for: %s\n\n", pageInput.Term)
	for i, mod := range mods {
		fmt.Printf("\t%02d ~ %s ~ v%s ~ by %s ~ %s\n", i+1, mod.Name(), mod.Version(), mod.Author(), page.trimString(mod.Description()))
	}

	fmt.Print("\n")
	fmt.Print("What to Do?\n")
	fmt.Print("-----------\n")
	fmt.Print("An) Add nth Mod\n")
	fmt.Print("S) Search New Term\n")
	fmt.Printf("T) Search Thunderstore for: %s\n", pageInput.Term)
	fmt.Print("Q) Cancel\n")
	fmt.Print("\n")

	return page.cachedResultsOptions(pageInput, mods)
}

func (page ModSearchResultsPage) cachedResultsOptions(pageInput input.ModSearchResultsPageInput, mods []mod.Mod) viewer.Options {
	addModArgs := make([]input.AddModTaskInput, len(mods))
	for i, mod := range mods {
		addModArgs[i] = input.AddModTaskInput{
			CachedMod:    mod,
			Profile:      pageInput.Profile,
			UseCachedMod: true,
		}
	}

	addMod := viewer.NewOption(viewer.OptionDto{
		Letter:   'A',
		Task:     viewer.TaskAddMod,
		TakesNum: true,
	}, addModArgs)

	searchNewTerm := viewer.NewOption(viewer.OptionDto{
		Letter: 'S',
		Task:   viewer.TaskSearchTerm,
	}, []input.SearchTermTaskInput{{Profile: pageInput.Profile}})

	searchThunderstore := viewer.NewOption(viewer.OptionDto{
		Letter: 'T',
		Page:   viewer.PageModSearchResults,
	}, []input.ModSearchResultsPageInput{
		{
			Profile:         pageInput.Profile,
			Term:            pageInput.Term,
			SkipCacheSearch: true,
		}})

	cancel := viewer.NewOption(viewer.OptionDto{
		Letter: 'Q',
		Page:   viewer.PageProfileViewer,
	}, []profile.Profile{pageInput.Profile})

	return viewer.NewOptions(
		[]viewer.Option{
			addMod,
			searchNewTerm,
			searchThunderstore,
			cancel,
		},
	)
}

func (page ModSearchResultsPage) showSearchResults(pageInput input.ModSearchResultsPageInput, results []mod.SearchResult) viewer.Options {
	clear()

	fmt.Printf("Adding Mod to %s\n", pageInput.Profile.Name())
	fmt.Print("---------------------------------------\n\n")

	fmt.Printf("Thunderstore Search Results for: %s\n\n", pageInput.Term)
	for i, result := range results {
		fmt.Printf("\t%02d ~ %s ~ by %s ~ %s\n", i+1, result.Name(), result.Author(), page.trimString(result.Description()))
	}

	fmt.Print("\n")
	fmt.Print("What to Do?\n")
	fmt.Print("-----------\n")
	fmt.Print("An) Add nth Mod\n")
	fmt.Print("S) Search New Term\n")
	fmt.Print("Q) Cancel\n")
	fmt.Print("\n")

	return page.thunderstoreResultsOptions(pageInput, results)
}

func (page ModSearchResultsPage) thunderstoreResultsOptions(pageInput input.ModSearchResultsPageInput, results []mod.SearchResult) viewer.Options {
	addModArgs := make([]input.AddModTaskInput, len(results))
	for i, result := range results {
		addModArgs[i] = input.AddModTaskInput{
			SearchResult: result,
			Profile:      pageInput.Profile,
			UseCachedMod: false,
		}
	}

	addMod := viewer.NewOption(viewer.OptionDto{
		Letter:   'A',
		Task:     viewer.TaskAddMod,
		TakesNum: true,
	}, addModArgs)

	searchNewTerm := viewer.NewOption(viewer.OptionDto{
		Letter: 'S',
		Task:   viewer.TaskSearchTerm,
	}, []input.SearchTermTaskInput{{Profile: pageInput.Profile}})

	cancel := viewer.NewOption(viewer.OptionDto{
		Letter: 'Q',
		Page:   viewer.PageProfileViewer,
	}, []profile.Profile{pageInput.Profile})

	return viewer.NewOptions(
		[]viewer.Option{
			addMod,
			searchNewTerm,
			cancel,
		},
	)
}

func (page ModSearchResultsPage) termHasSpaces(term string) string {
	i := strings.LastIndex(term, " ")

	if i > 0 && i < len(term) {
		return "not"
	}

	return ""
}

func (page ModSearchResultsPage) trimString(s string) string {
	cutoff := 120

	if len(s) < cutoff {
		return s
	}

	return s[:cutoff] + "..."
}
