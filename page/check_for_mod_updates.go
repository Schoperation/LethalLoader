package page

import (
	"fmt"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type latestModListingsGetter interface {
	GetByNameAndAuthor(name, author string) (mod.Listing, error)
}

type CheckForModUpdatesPage struct {
	modListingGetter latestModListingsGetter
}

func NewCheckForModUpdatesPage(
	modListingGetter latestModListingsGetter,
) CheckForModUpdatesPage {
	return CheckForModUpdatesPage{
		modListingGetter: modListingGetter,
	}
}

func (page CheckForModUpdatesPage) Show(args any) (viewer.OptionsResult, error) {
	clear()

	pf, ok := args.(profile.Profile)
	if !ok {
		return viewer.OptionsResult{}, fmt.Errorf("could not cast profile")
	}

	fmt.Print("Checking for updates...\n")

	type outdatedMod struct {
		Current mod.Mod
		Latest  mod.Listing
	}

	var outdatedMods []outdatedMod

	for _, mod := range pf.Mods() {
		latestListing, err := page.modListingGetter.GetByNameAndAuthor(mod.Name(), mod.Author())
		if err != nil {
			return viewer.OptionsResult{}, err
		}

		if latestListing.DateCreated().After(mod.DateCreated()) {
			outdatedMods = append(outdatedMods, outdatedMod{
				Current: mod,
				Latest:  latestListing,
			})
		}
	}

	clear()

	fmt.Printf("Updating Mods for Profile %s\n", pf.Name())
	fmt.Print("---------------------------------------\n\n")

	if len(outdatedMods) == 0 {
		fmt.Print("\tAll mods are up to date!\n")
	}

	for i, mod := range outdatedMods {
		fmt.Printf("\t%02d ~ %s ~ from v%s to v%s ~ updated %s\n", i+1, mod.Current.Name(), mod.Current.Version(), mod.Latest.Version(), mod.Latest.DateCreated().Format("01/02/2006 15:04"))
	}

	fmt.Print("\n")
	fmt.Print("What to Do?\n")
	fmt.Print("-----------\n")

	if len(outdatedMods) > 0 {
		fmt.Print("A ) Update All\n")
		fmt.Print("Un) Update nth Mod\n")
	}

	fmt.Printf("Q ) Back to Profile %s\n", pf.Name())
	fmt.Print("\n")

	latestListings := make([]mod.Listing, len(outdatedMods))
	for i, mod := range outdatedMods {
		latestListings[i] = mod.Latest
	}

	options := page.options(pf, latestListings)
	return options.TakeInput(), nil
}

func (page CheckForModUpdatesPage) options(pf profile.Profile, latestListings []mod.Listing) viewer.Options {
	updateAll := viewer.NewOption(viewer.OptionDto{
		Letter: 'A',
		Task:   viewer.TaskUpdateMods,
	}, []input.UpdateModsTaskInput{
		{
			Listings: latestListings,
			Profile:  pf,
		},
	})

	updateNthArgs := make([]input.UpdateModsTaskInput, len(latestListings))
	for i, listing := range latestListings {
		updateNthArgs[i] = input.UpdateModsTaskInput{
			Listings: []mod.Listing{listing},
			Profile:  pf,
		}
	}

	updateNth := viewer.NewOption(viewer.OptionDto{
		Letter:   'U',
		Task:     viewer.TaskUpdateMods,
		TakesNum: true,
	}, updateNthArgs)

	back := viewer.NewOption(viewer.OptionDto{
		Letter: 'Q',
		Page:   viewer.PageProfileViewer,
	}, []profile.Profile{pf})

	if len(latestListings) == 0 {
		return viewer.NewOptions([]viewer.Option{back})
	}

	return viewer.NewOptions([]viewer.Option{
		updateAll,
		updateNth,
		back,
	})
}
