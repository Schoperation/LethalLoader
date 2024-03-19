package task

import (
	"fmt"
	"schoperation/lethalloader/domain/config"
	"schoperation/lethalloader/domain/input"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type selectedProfileConfigChanger interface {
	Get() (config.MainConfig, error)
	Save(mainConfig config.MainConfig) error
}

type profileSwitcher interface {
	Switch(oldPf profile.Profile, newPf profile.Profile, gameFilesPath string) error
}

type SwitchProfileTask struct {
	configChanger   selectedProfileConfigChanger
	profileSwitcher profileSwitcher
}

func NewSwitchProfileTask(
	configChanger selectedProfileConfigChanger,
	profileSwitcher profileSwitcher,
) SwitchProfileTask {
	return SwitchProfileTask{
		configChanger:   configChanger,
		profileSwitcher: profileSwitcher,
	}
}

func (task SwitchProfileTask) Do(args any) (viewer.TaskResult, error) {
	taskInput, ok := args.(input.SwitchProfileTaskInput)
	if !ok {
		return viewer.TaskResult{}, fmt.Errorf("could not cast input")
	}

	mainConfig, err := task.configChanger.Get()
	if err != nil {
		return viewer.TaskResult{}, err
	}

	err = task.profileSwitcher.Switch(taskInput.OldProfile, taskInput.NewProfile, mainConfig.GameFilesPath())
	if err != nil {
		return viewer.TaskResult{}, err
	}

	mainConfig.UpdateSelectedProfile(taskInput.NewProfile.Name())
	err = task.configChanger.Save(mainConfig)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	return viewer.NewTaskResult(viewer.PageMainMenu, nil), nil
}
