package main

import (
	"fmt"
	"schoperation/lethalloader/domain/option"
)

type cliPage interface {
	Show(args ...any) (option.OptionsResults, error)
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
	for {
		taskResult, err := (*viewer.currentTask).Do()
		if err != nil {
			return err
		}

		results, err := (*viewer.currentPage).Show(taskResult)
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
