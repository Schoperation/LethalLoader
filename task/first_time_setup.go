package task

import (
	"fmt"
	"schoperation/lethalloader/domain/config"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
	"strings"
)

type mainConfigCreator interface {
	Get() (config.MainConfig, error)
	Save(mainConfig config.MainConfig) error
}

type gameFilesPathChecker interface {
	CheckDefaultPath() (string, error)
	CheckPath(path string) (bool, error)
}

type vanillaProfileSaver interface {
	Save(pf profile.Profile) error
}

type FirstTimeSetupTask struct {
	mainConfigCreator    mainConfigCreator
	gameFilesPathChecker gameFilesPathChecker
	vanillaProfileSaver  vanillaProfileSaver
}

func NewFirstTimeSetupTask(
	mainConfigCreator mainConfigCreator,
	gameFilesPathChecker gameFilesPathChecker,
	vanillaProfileSaver vanillaProfileSaver,
) FirstTimeSetupTask {
	return FirstTimeSetupTask{
		mainConfigCreator:    mainConfigCreator,
		gameFilesPathChecker: gameFilesPathChecker,
		vanillaProfileSaver:  vanillaProfileSaver,
	}
}

func (task FirstTimeSetupTask) Do(args any) (viewer.TaskResult, error) {
	mainConfig, err := task.mainConfigCreator.Get()
	if err != nil {
		return viewer.TaskResult{}, err
	}

	if mainConfig.GameFilesPath() != "" {
		return viewer.NewTaskResult(viewer.PageMainMenu, nil), nil
	}

	fmt.Printf("First time? Trying to find your Lethal Company game files...\n")

	gameFilePath, err := task.gameFilesPathChecker.CheckDefaultPath()
	if err != nil {
		return viewer.TaskResult{}, err
	}

	gameFilePath = strings.TrimSpace(gameFilePath)

	if gameFilePath == "" {
		fmt.Printf("Couldn't find your game files. Would you be so polite to tell us where they are?\n")

		gameFilePath, err = task.customGameFilePath()
		if err != nil {
			return viewer.TaskResult{}, err
		}
	} else {
		fmt.Printf("Found existing game files: %s\n", gameFilePath)
		fmt.Printf("Are we good? (Y/n)\n")
		fmt.Printf("\n")

		for {
			var weGood string
			fmt.Print(">")
			fmt.Scanf("%s", &weGood)

			if weGood == "" || weGood == "y" {
				break
			}

			if weGood == "n" {
				gameFilePath, err = task.customGameFilePath()
				if err != nil {
					return viewer.TaskResult{}, err
				}
				break
			}

			fmt.Printf("Bruh what you saying\n")
		}
	}

	mainConfig.UpdateGameFilesPath(gameFilePath)

	vanillaProfile, err := profile.NewBlankProfile(profile.ProfileDto{
		Name: "Vanilla",
	})
	if err != nil {
		return viewer.TaskResult{}, err
	}

	err = task.vanillaProfileSaver.Save(vanillaProfile)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	mainConfig.UpdateSelectedProfile(vanillaProfile.Name())

	err = task.mainConfigCreator.Save(mainConfig)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	return viewer.NewTaskResult(viewer.PageMainMenu, nil), nil
}

func (task FirstTimeSetupTask) customGameFilePath() (string, error) {
	fmt.Printf("Type the full path (C:\\Program Files (x86)\\whatever)\n")

	gameFilePath := ""
	for {
		var path string
		fmt.Scanf("%s", &path)

		exists, err := task.gameFilesPathChecker.CheckPath(path)
		if err != nil {
			return "", err
		}

		if exists {
			gameFilePath = path
			break
		}

		fmt.Printf("Bruh that don't exist\n")
	}

	return gameFilePath, nil
}
