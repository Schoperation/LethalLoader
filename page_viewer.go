package main

import (
	"schoperation/lethalloader/domain/option"
)

type currentPage interface {
	Show() (option.OptionsResults, error)
}

type currentTask interface {
	Do(args any) error
}

type PageViewer struct {
	currentPage currentPage
	pages       map[option.CmdName]currentPage
}

func NewPageViewer(
	mainMenuPage currentPage,
) PageViewer {
	pages := map[option.CmdName]currentPage{
		option.PageMainMenu: mainMenuPage,
	}

	return PageViewer{
		currentPage: mainMenuPage,
		pages:       pages,
	}
}

func (viewer PageViewer) Run() error {
	for {
		results, err := viewer.currentPage.Show()
		if err != nil {
			return err
		}

		if results.CmdName.IsQuit() {
			return nil
		}

		if results.CmdName.IsPage() {
			viewer.currentPage = viewer.pages[results.CmdName]
			continue
		}
	}
}
