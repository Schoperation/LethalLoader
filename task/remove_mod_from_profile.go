package task

import (
	"fmt"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type profileWithRemovedModSaver interface {
	Save(pf profile.Profile) error
}

type RemoveModFromProfileTask struct {
	profileSaver profileWithRemovedModSaver
}

func NewRemoveModFromProfileTask(
	profileSaver profileWithRemovedModSaver,
) RemoveModFromProfileTask {
	return RemoveModFromProfileTask{
		profileSaver: profileSaver,
	}
}

func (task RemoveModFromProfileTask) Do(args any) (viewer.TaskResult, error) {
	taskInput, ok := args.(input.RemoveModFromProfileTaskInput)
	if !ok {
		return viewer.TaskResult{}, fmt.Errorf("could not cast input")
	}

	// TODO add confirmation if there are mods that depend on it

	taskInput.Profile.RemoveMod(taskInput.Mod)

	err := task.profileSaver.Save(taskInput.Profile)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	return viewer.NewTaskResult(viewer.PageProfileViewer, taskInput.Profile), nil
}
