package task

import (
	"fmt"
	"schoperation/lethalloader/domain/config"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
	"time"
)

type mainConfigGetter interface {
	Get() (config.MainConfig, error)
}

type profileDeleter interface {
	Delete(pf profile.Profile) error
}

type DeleteProfileTask struct {
	mainConfigGetter mainConfigGetter
	profileDeleter   profileDeleter
}

func NewDeleteProfileTask(
	mainConfigGetter mainConfigGetter,
	profileDeleter profileDeleter,
) DeleteProfileTask {
	return DeleteProfileTask{
		mainConfigGetter: mainConfigGetter,
		profileDeleter:   profileDeleter,
	}
}

func (task DeleteProfileTask) Do(args any) (viewer.TaskResult, error) {
	profile, ok := args.(profile.Profile)
	if !ok {
		return viewer.TaskResult{}, fmt.Errorf("could not cast profile")
	}

	mainConfig, err := task.mainConfigGetter.Get()
	if err != nil {
		return viewer.TaskResult{}, err
	}

	if mainConfig.SelectedProfile() == profile.Name() {
		fmt.Printf("Cannot delete: profile selected. Please select another profile before deleting.\n")
		time.Sleep(3 * time.Second)
		return viewer.NewTaskResult(viewer.PageMainMenu, nil), nil
	}

	err = task.profileDeleter.Delete(profile)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	return viewer.NewTaskResult(viewer.PageMainMenu, nil), nil
}
