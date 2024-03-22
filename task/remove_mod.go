package task

import (
	"fmt"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type profileWithRemovedModSaver interface {
	Save(pf profile.Profile) error
}

type gameFilesRemover interface {
	DeleteMod(mod mod.Mod, pfName string) error
}

type RemoveModTask struct {
	profileSaver     profileWithRemovedModSaver
	gameFilesRemover gameFilesRemover
}

func NewRemoveModTask(
	profileSaver profileWithRemovedModSaver,
	gameFilesRemover gameFilesRemover,
) RemoveModTask {
	return RemoveModTask{
		profileSaver:     profileSaver,
		gameFilesRemover: gameFilesRemover,
	}
}

func (task RemoveModTask) Do(args any) (viewer.TaskResult, error) {
	taskInput, ok := args.(input.RemoveModTaskInput)
	if !ok {
		return viewer.TaskResult{}, fmt.Errorf("could not cast input")
	}

	// TODO add confirmation if there are mods that depend on it

	taskInput.Profile.RemoveMod(taskInput.Mod)

	err := task.gameFilesRemover.DeleteMod(taskInput.Mod, taskInput.Profile.Name())
	if err != nil {
		return viewer.TaskResult{}, err
	}

	err = task.profileSaver.Save(taskInput.Profile)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	return viewer.NewTaskResult(viewer.PageProfileViewer, taskInput.Profile), nil
}
