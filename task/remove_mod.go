package task

import (
	"fmt"
	"schoperation/lethalloader/domain/config"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type profileWithRemovedModSaver interface {
	Switch(oldPf profile.Profile, newPf profile.Profile, gameFilesPath string) error
	Save(pf profile.Profile) error
}

type removeModConfigGetter interface {
	Get() (config.MainConfig, error)
}

type RemoveModTask struct {
	profileSaver profileWithRemovedModSaver
	configGetter removeModConfigGetter
}

func NewRemoveModTask(
	profileSaver profileWithRemovedModSaver,
	configGetter removeModConfigGetter,
) RemoveModTask {
	return RemoveModTask{
		profileSaver: profileSaver,
		configGetter: configGetter,
	}
}

func (task RemoveModTask) Do(args any) (viewer.TaskResult, error) {
	taskInput, ok := args.(input.RemoveModTaskInput)
	if !ok {
		return viewer.TaskResult{}, fmt.Errorf("could not cast input")
	}

	// TODO add confirmation if there are mods that depend on it

	taskInput.Profile.RemoveMod(taskInput.Mod)

	mainConfig, err := task.configGetter.Get()
	if err != nil {
		return viewer.TaskResult{}, err
	}

	if mainConfig.SelectedProfile() == taskInput.Profile.Name() {
		err = task.profileSaver.Switch(taskInput.Profile, taskInput.Profile, mainConfig.GameFilesPath())
		if err != nil {
			return viewer.TaskResult{}, err
		}
	}

	err = task.profileSaver.Save(taskInput.Profile)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	return viewer.NewTaskResult(viewer.PageProfileViewer, taskInput.Profile), nil
}
