package task

import (
	"fmt"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/viewer"
)

type latestModGetter interface {
	GetByModListing(listing mod.Listing) (mod.Mod, error)
}

type UpdateModsTask struct {
	modGetter latestModGetter
}

func NewUpdateModsTask(
	modGetter latestModGetter,
) UpdateModsTask {
	return UpdateModsTask{
		modGetter: modGetter,
	}
}

func (task UpdateModsTask) Do(args any) (viewer.TaskResult, error) {
	taskInput, ok := args.(input.UpdateModsTaskInput)
	if !ok {
		return viewer.TaskResult{}, fmt.Errorf("could not parse input")
	}

	// TODO
}
