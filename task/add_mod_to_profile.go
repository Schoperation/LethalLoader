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
}

type searchedModDownloader interface {
	GetByModListing(listing mod.Listing) (mod.Mod, error)
}

type profileSaver interface {
	Save(pf profile.Profile) error
}

type AddModToProfileTask struct {
	searchedModListingGetter searchedModListingGetter
	searchedModDownloader    searchedModDownloader
	profileSaver             profileSaver
}

func NewAddModToProfileTask(
	searchedModListingGetter searchedModListingGetter,
	searchedModDownloader searchedModDownloader,
	profileSaver profileSaver,
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

	// TODO add deps (other than bepinex)

	err := taskInput.Profile.AddMod(newMod)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	err = task.profileSaver.Save(taskInput.Profile)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	return viewer.NewTaskResult(viewer.PageProfileViewer, taskInput.Profile), nil
}
