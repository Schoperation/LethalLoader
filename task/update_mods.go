package task

import (
	"fmt"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type latestModGetter interface {
	GetByModListing(listing mod.Listing) (mod.Mod, error)
}

type gameFilesUpdater interface {
	DeleteMod(mod mod.Mod, pfName string) error
	AddMod(mod mod.Mod, pfName string) error
}

type updatedProfileSaver interface {
	Save(pf profile.Profile) error
}

type UpdateModsTask struct {
	modGetter        latestModGetter
	gameFilesUpdater gameFilesUpdater
	profileSaver     updatedProfileSaver
}

func NewUpdateModsTask(
	modGetter latestModGetter,
	gameFilesUpdater gameFilesUpdater,
	profileSaver updatedProfileSaver,
) UpdateModsTask {
	return UpdateModsTask{
		modGetter:        modGetter,
		gameFilesUpdater: gameFilesUpdater,
		profileSaver:     profileSaver,
	}
}

func (task UpdateModsTask) Do(args any) (viewer.TaskResult, error) {
	taskInput, ok := args.(input.UpdateModsTaskInput)
	if !ok {
		return viewer.TaskResult{}, fmt.Errorf("could not parse input")
	}

	for _, listing := range taskInput.Listings {
		oldMod, err := taskInput.Profile.Mod(listing.Name())
		if err != nil {
			return viewer.TaskResult{}, err
		}

		fmt.Printf("Updating %s from %s to %s...\n", listing.Name(), oldMod.Version(), listing.Version())

		updatedMod, err := task.modGetter.GetByModListing(listing)
		if err != nil {
			return viewer.TaskResult{}, err
		}

		err = task.gameFilesUpdater.DeleteMod(oldMod, taskInput.Profile.Name())
		if err != nil {
			return viewer.TaskResult{}, err
		}

		taskInput.Profile.RemoveMod(oldMod)
		taskInput.Profile.AddMod(updatedMod)

		err = task.gameFilesUpdater.AddMod(updatedMod, taskInput.Profile.Name())
		if err != nil {
			return viewer.TaskResult{}, err
		}
	}

	err := task.profileSaver.Save(taskInput.Profile)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	return viewer.NewTaskResult(viewer.PageCheckForModUpdates, taskInput.Profile), nil
}
