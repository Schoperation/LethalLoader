package viewer

import (
	"fmt"
	"schoperation/lethalloader/domain/viewer"
)

type cliPage interface {
	Show(args ...any) (viewer.OptionsResult, error)
}

type cliTask interface {
	Do(args ...any) (any, error)
}

type PageViewer struct {
	currentTask *cliTask
	currentPage *cliPage
	tasks       map[viewer.Task]cliTask
	pages       map[viewer.Page]cliPage
}

func NewPageViewer(
	mainMenuPage cliPage,
	profileViewerPage cliPage,
	firstTimeSetupTask cliTask,
	newProfileTask cliTask,
) PageViewer {
	tasks := map[viewer.Task]cliTask{
		viewer.TaskFirstTimeSetup: firstTimeSetupTask,
		viewer.TaskNewProfile:     newProfileTask,
	}

	pages := map[viewer.Page]cliPage{
		viewer.PageMainMenu:      mainMenuPage,
		viewer.PageProfileViewer: profileViewerPage,
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
