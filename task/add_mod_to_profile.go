package task

import (
	"fmt"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type searchedModListingGetter interface {
	GetByNameAndAuthor(name, author string) (mod.Listing, error)
	GetBySlug(slug mod.Slug) (mod.Listing, error)
}

type searchedModDownloader interface {
	GetByModListing(listing mod.Listing) (mod.Mod, error)
}

type profileWithNewModSaver interface {
	Save(pf profile.Profile) error
}

type AddModToProfileTask struct {
	searchedModListingGetter searchedModListingGetter
	searchedModDownloader    searchedModDownloader
	profileSaver             profileWithNewModSaver
}

func NewAddModToProfileTask(
	searchedModListingGetter searchedModListingGetter,
	searchedModDownloader searchedModDownloader,
	profileSaver profileWithNewModSaver,
) AddModToProfileTask {
	return AddModToProfileTask{
		searchedModListingGetter: searchedModListingGetter,
		searchedModDownloader:    searchedModDownloader,
		profileSaver:             profileSaver,
	}
}

func (task AddModToProfileTask) Do(args any) (viewer.TaskResult, error) {
	taskInput, ok := args.(input.AddModToProfileTaskInput)
	if !ok {
		return viewer.TaskResult{}, fmt.Errorf("could not cast input")
	}

	modName := taskInput.CachedMod.Name()
	if !taskInput.UseCachedMod {
		modName = taskInput.SearchResult.Name()
	}

	fmt.Printf("Adding %s...\n", modName)

	newMod := taskInput.CachedMod
	if !taskInput.UseCachedMod {
		listing, err := task.searchedModListingGetter.GetByNameAndAuthor(taskInput.SearchResult.Name(), taskInput.SearchResult.Author())
		if err != nil {
			return viewer.TaskResult{}, err
		}

		newMod, err = task.searchedModDownloader.GetByModListing(listing)
		if err != nil {
			return viewer.TaskResult{}, err
		}
	}

	taskInput.Profile.AddMod(newMod)

	additionalMods, err := task.addDependencies(newMod)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	for _, addMod := range additionalMods {
		taskInput.Profile.AddMod(addMod)
	}

	err = task.profileSaver.Save(taskInput.Profile)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	return viewer.NewTaskResult(viewer.PageProfileViewer, taskInput.Profile), nil
}

func (task AddModToProfileTask) addDependencies(newMod mod.Mod) ([]mod.Mod, error) {
	if len(newMod.Dependencies()) == 0 {
		return nil, nil
	}

	pluralDeps := "dependencies"
	if len(newMod.Dependencies()) == 1 {
		pluralDeps = "dependency"
	}

	fmt.Printf("Adding %d %s for %s...\n", len(newMod.Dependencies()), pluralDeps, newMod.Name())
	var additionalMods []mod.Mod

	for _, dep := range newMod.Dependencies() {
		listing, err := task.searchedModListingGetter.GetBySlug(dep)
		if err != nil {
			return nil, err
		}

		depMod, err := task.searchedModDownloader.GetByModListing(listing)
		if err != nil {
			return nil, err
		}

		additionalMods = append(additionalMods, depMod)

		moreDeps, err := task.addDependencies(depMod)
		if err != nil {
			return nil, err
		}

		additionalMods = append(additionalMods, moreDeps...)
	}

	return additionalMods, nil
}
