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

type updatedProfileSaver interface {
	Save(pf profile.Profile) error
}

type UpdateModsTask struct {
	modGetter    latestModGetter
	profileSaver updatedProfileSaver
}

func NewUpdateModsTask(
	modGetter latestModGetter,
	profileSaver updatedProfileSaver,
) UpdateModsTask {
	return UpdateModsTask{
		modGetter:    modGetter,
		profileSaver: profileSaver,
	}
}

func (task UpdateModsTask) Do(args any) (viewer.TaskResult, error) {
	taskInput, ok := args.(input.UpdateModsTaskInput)
	if !ok {
		return viewer.TaskResult{}, fmt.Errorf("could not parse input")
	}

	for _, listing := range taskInput.Listings {
		fmt.Printf("Updating %s to %s...\n", listing.Name(), listing.Version())

		updatedMod, err := task.modGetter.GetByModListing(listing)
		if err != nil {
			return viewer.TaskResult{}, err
		}

		taskInput.Profile.RemoveMod(updatedMod)
		taskInput.Profile.AddMod(updatedMod)
	}

	err := task.profileSaver.Save(taskInput.Profile)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	return viewer.NewTaskResult(viewer.PageCheckForModUpdates, taskInput.Profile), nil
}
