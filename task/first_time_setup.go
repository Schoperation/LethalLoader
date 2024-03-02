package task

import (
	"fmt"
	"schoperation/lethalloader/domain/config"
	"schoperation/lethalloader/domain/profile"
)

type mainConfigCreator interface {
	Get() (config.MainConfig, error)
	Save(mainConfig config.MainConfig) error
}

type steamChecker interface {
	CheckDefault() (string, error)
	Check(path string) (bool, error)
}

type vanillaProfileSaver interface {
	Save(pf profile.Profile) error
}

type FirstTimeSetupTask struct {
	mainConfigCreator   mainConfigCreator
	steamChecker        steamChecker
	vanillaProfileSaver vanillaProfileSaver
}

func NewFirstTimeSetupTask(
	mainConfigCreator mainConfigCreator,
	steamChecker steamChecker,
	vanillaProfileSaver vanillaProfileSaver,
) FirstTimeSetupTask {
	return FirstTimeSetupTask{
		mainConfigCreator:   mainConfigCreator,
		steamChecker:        steamChecker,
		vanillaProfileSaver: vanillaProfileSaver,
	}
}

func (task FirstTimeSetupTask) Do(args ...any) (any, error) {
	mainConfig, err := task.mainConfigCreator.Get()
	if err != nil {
		return nil, err
	}

	if mainConfig.GameFilePath() != "" {
		return nil, nil
	}

	fmt.Printf("First time? Trying to find your Lethal Company game files...\n")

	gameFilePath, err := task.steamChecker.CheckDefault()
	if err != nil {
		return nil, err
	}

	if gameFilePath == "" {
		fmt.Printf("Couldn't find your game files. Would you be so polite to tell us where they are?\n")
		fmt.Printf("Type the full path (C:\\Program Files (x86)\\whatever)\n")

		for {
			fmt.Scanf("%s", &gameFilePath)

			exists, err := task.steamChecker.Check(gameFilePath)
			if err != nil {
				return nil, err
			}

			if exists {
				break
			}

			fmt.Printf("Bruh that don't exist\n")
		}
	}

	mainConfig.UpdateGameFilePath(gameFilePath)

	vanillaProfile, err := profile.NewBlankProfile(profile.ProfileDto{
		Name: "Vanilla",
	})
	if err != nil {
		return nil, err
	}

	err = task.vanillaProfileSaver.Save(vanillaProfile)
	if err != nil {
		return nil, err
	}

	mainConfig.UpdateSelectedProfile(vanillaProfile.Name())

	err = task.mainConfigCreator.Save(mainConfig)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
