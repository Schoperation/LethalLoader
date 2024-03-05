package viewer

import (
	"fmt"
	"schoperation/lethalloader/domain/option"
)

type cliPage interface {
	Show(args ...any) (option.Options, error)
}

type cliTask interface {
	Do(args ...any) (any, error)
}

type PageViewer struct {
	currentTask *cliTask
	currentPage *cliPage
	tasks       map[option.TaskName]cliTask
	pages       map[option.PageName]cliPage
}

func NewPageViewer(
	mainMenuPage cliPage,
	profileViewerPage cliPage,
	firstTimeSetupTask cliTask,
	newProfileTask cliTask,
) PageViewer {
	tasks := map[option.TaskName]cliTask{
		option.TaskFirstTimeSetup: firstTimeSetupTask,
		option.TaskNewProfile:     newProfileTask,
	}

	pages := map[option.PageName]cliPage{
		option.PageMainMenu:      mainMenuPage,
		option.PageProfileViewer: profileViewerPage,
	}

	return PageViewer{
		currentPage: &mainMenuPage,
		currentTask: &firstTimeSetupTask,
		pages:       pages,
		tasks:       tasks,
	}
}

func (viewer PageViewer) Run() error {
	var args any
	var err error

	options := viewer.NewOption(viewer.OptionDto{}, []string{})

	for {
		args, err = viewer.doTask()
		if err != nil {
			return err
		}

		options, err := viewer.showPage(args)
		if err != nil {
			return err
		}

		if results.Task == option.TaskQuit {
			return nil
		}

		viewer.currentTask = nil
		if results.Task != "" {
			nextTask, exists := viewer.tasks[results.Task]
			if !exists {
				return fmt.Errorf("could not find task %s", results.Task)
			}

			viewer.currentTask = &nextTask
		}

		viewer.currentPage = nil
		if results.Page != "" {
			nextPage, exists := viewer.pages[results.Page]
			if !exists {
				return fmt.Errorf("could not find page %s", results.Page)
			}

			viewer.currentPage = &nextPage
		}
	}
}

func (viewer PageViewer) doTask() (any, error) {
	if viewer.currentTask == nil {
		return nil, nil
	}

	return (*viewer.currentTask).Do()
}

func (viewer PageViewer) showPage(args ...any) (option.Options, error) {
	if viewer.currentPage == nil {
		return option.Options{}, nil
	}

	return (*viewer.currentPage).Show(args...)
}
