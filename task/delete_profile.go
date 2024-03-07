package task

import (
	"fmt"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type profileDeleter interface {
	Delete(pf profile.Profile) error
}

type DeleteProfileTask struct {
	profileDeleter profileDeleter
}

func NewDeleteProfileTask(
	profileDeleter profileDeleter,
) DeleteProfileTask {
	return DeleteProfileTask{
		profileDeleter: profileDeleter,
	}
}

func (task DeleteProfileTask) Do(args any) (viewer.TaskResult, error) {
	profile, ok := args.(profile.Profile)
	if !ok {
		return viewer.TaskResult{}, fmt.Errorf("could not cast profile")
	}

	err := task.profileDeleter.Delete(profile)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	return viewer.NewTaskResult(viewer.PageMainMenu, nil), nil
}
