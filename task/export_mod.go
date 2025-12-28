package task

import (
	"fmt"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
	"time"
)

type profileExporter interface {
	Export(pf profile.Profile) (string, error)
}

type ExportProfileTask struct {
	profileExporter profileExporter
}

func NewExportProfileTask(
	profileExporter profileExporter,
) ExportProfileTask {
	return ExportProfileTask{
		profileExporter: profileExporter,
	}
}

func (task ExportProfileTask) Do(args any) (viewer.TaskResult, error) {
	profile, ok := args.(profile.Profile)
	if !ok {
		return viewer.TaskResult{}, fmt.Errorf("could not cast profile")
	}

	fileName, err := task.profileExporter.Export(profile)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	fmt.Printf("Exported into %s.\n", fileName)
	time.Sleep(3 * time.Second)

	return viewer.NewTaskResult(viewer.PageProfileViewer, profile), nil
}
