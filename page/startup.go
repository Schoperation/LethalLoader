package page

import (
	"fmt"
	"schoperation/lethalloader/domain/config"
	"schoperation/lethalloader/domain/viewer"
)

type mainConfigGetter interface {
	Get() (config.MainConfig, error)
}

type StartupPage struct {
	mainConfigGetter mainConfigGetter
}

func NewStartupPage(
	mainConfigGetter mainConfigGetter,
) StartupPage {
	return StartupPage{
		mainConfigGetter: mainConfigGetter,
	}
}

func (page StartupPage) Show(args any) (viewer.OptionsResult, error) {
	fmt.Printf("Starting LethalLoader...\n")

	mainConfig, err := page.mainConfigGetter.Get()
	if err != nil {
		return viewer.OptionsResult{}, err
	}

	if mainConfig.GameFilesPath() != "" {
		return viewer.NewJumpToPageOptionsResult(viewer.PageMainMenu), nil
	}

	return viewer.NewJumpToTaskOptionsResult(viewer.TaskFirstTimeSetup), nil
}
