package viewer

import (
	"fmt"
	"schoperation/lethalloader/domain/viewer"
)

type cliPage interface {
	Show(args any) (viewer.OptionsResult, error)
}

type cliTask interface {
	Do(args any) (viewer.TaskResult, error)
}

type PageViewer struct {
	tasks map[viewer.Task]cliTask
	pages map[viewer.Page]cliPage
}

func NewPageViewer(
	mainMenuPage cliPage,
	profileViewerPage cliPage,
	modSearchResultsPage cliPage,
	firstTimeSetupTask cliTask,
	newProfileTask cliTask,
	deleteProfileTask cliTask,
	searchTermTask cliTask,
	addModToProfileTask cliTask,
) PageViewer {
	tasks := map[viewer.Task]cliTask{
		viewer.TaskFirstTimeSetup:  firstTimeSetupTask,
		viewer.TaskNewProfile:      newProfileTask,
		viewer.TaskDeleteProfile:   deleteProfileTask,
		viewer.TaskSearchTerm:      searchTermTask,
		viewer.TaskAddModToProfile: addModToProfileTask,
	}

	pages := map[viewer.Page]cliPage{
		viewer.PageMainMenu:         mainMenuPage,
		viewer.PageProfileViewer:    profileViewerPage,
		viewer.PageModSearchResults: modSearchResultsPage,
	}

	return PageViewer{
		pages: pages,
		tasks: tasks,
	}
}

func (view PageViewer) Run() error {
	currentTask := viewer.TaskFirstTimeSetup
	var currentPage viewer.Page
	var args any
	var err error

	for {
		if currentTask == "" && currentPage == "" {
			return fmt.Errorf("no task or page selected! bruh")
		}

		if currentTask != "" {
			if currentTask == viewer.TaskQuit {
				return nil
			}

			currentPage, args, err = view.doTask(currentTask, args)
			if err != nil {
				return err
			}
		}

		if currentPage != "" {
			currentTask, currentPage, args, err = view.showPage(currentPage, args)
			if err != nil {
				return err
			}
		}
	}
}

func (view PageViewer) doTask(task viewer.Task, args any) (viewer.Page, any, error) {
	cliTask, exists := view.tasks[task]
	if !exists {
		return "", nil, fmt.Errorf("could not find task %s", task)
	}

	taskResults, err := cliTask.Do(args)
	if err != nil {
		return "", nil, err
	}

	return taskResults.NextPage(), taskResults.NextArgs(), nil
}

func (view PageViewer) showPage(page viewer.Page, args any) (viewer.Task, viewer.Page, any, error) {
	cliPage, exists := view.pages[page]
	if !exists {
		return "", "", nil, fmt.Errorf("could not find page %s", page)
	}

	optionsResult, err := cliPage.Show(args)
	if err != nil {
		return "", "", nil, err
	}

	return optionsResult.NextTask(), optionsResult.NextPage(), optionsResult.NextArgs(), nil
}
