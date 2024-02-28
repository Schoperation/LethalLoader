package main

import (
	"schoperation/lethalloader/domain/option"
)

type cliPage interface {
	Show(args ...any) (option.OptionsResults, error)
}

type cliTask interface {
	Do(args ...any) error
}

type PageViewer struct {
	currentTask cliTask
	currentPage cliPage
	tasks       map[option.TaskName]cliTask
	pages       map[option.PageName]cliPage
}

func NewPageViewer(
	mainMenuPage cliPage,
	firstTimeSetupTask cliTask,
) PageViewer {
	tasks := map[option.TaskName]cliTask{
		option.TaskFirstTimeSetup: firstTimeSetupTask,
	}

	pages := map[option.PageName]cliPage{
		option.PageMainMenu: mainMenuPage,
	}

	return PageViewer{
		currentPage: mainMenuPage,
		currentTask: firstTimeSetupTask,
		pages:       pages,
		tasks:       tasks,
	}
}

func (viewer PageViewer) Run() error {
	for {
		err := viewer.currentTask.Do()
		if err != nil {
			return err
		}

		results, err := viewer.currentPage.Show()
		if err != nil {
			return err
		}

		if results.Task == "quit" {
			return nil
		}

		if results.Task != "" {
			viewer.currentTask = viewer.tasks[results.Task]
		}

		if results.Page != "" {
			viewer.currentPage = viewer.pages[results.Page]
		}
	}
}
