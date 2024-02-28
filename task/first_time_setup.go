package task

import (
	"fmt"
	"schoperation/lethalloader/domain/config"
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
)

type mainConfigCreator interface {
	Read() (config.MainConfig, error)
	Write(mainConfig config.MainConfig) error
}

type steamChecker interface {
	CheckDefault() (string, error)
	Check(path string) (bool, error)
}

type vanillaProfileSaver interface {
	Save(pf profile.Profile) error
}

type modListCreator interface {
	SaveAllToList(mods []mod.Mod) error
}

type FirstTimeSetupTask struct {
	mainConfigCreator   mainConfigCreator
	steamChecker        steamChecker
	vanillaProfileSaver vanillaProfileSaver
	modListCreator      modListCreator
}

func NewFirstTimeSetupTask(
	mainConfigCreator mainConfigCreator,
	steamChecker steamChecker,
	vanillaProfileSaver vanillaProfileSaver,
	modListCreator modListCreator,
) FirstTimeSetupTask {
	return FirstTimeSetupTask{
		mainConfigCreator:   mainConfigCreator,
		steamChecker:        steamChecker,
		vanillaProfileSaver: vanillaProfileSaver,
		modListCreator:      modListCreator,
	}
}

func (task FirstTimeSetupTask) Do(args ...any) error {
	mainConfig, err := task.mainConfigCreator.Read()
	if err != nil {
		return err
	}

	if mainConfig.GameFilePath() != "" {
		return nil
	}

	fmt.Printf("First time? Trying to find your Lethal Company game files...\n")

	gameFilePath, err := task.steamChecker.CheckDefault()
	if err != nil {
		return err
	}

	if gameFilePath == "" {
		fmt.Printf("Couldn't find your game files. Would you be so polite to tell us where they are?\n")
		fmt.Printf("Type the full path (C:\\Program Files (x86)\\whatever)\n")

		for {
			fmt.Scanf("%s", &gameFilePath)

			exists, err := task.steamChecker.Check(gameFilePath)
			if err != nil {
				return err
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
		return err
	}

	err = task.vanillaProfileSaver.Save(vanillaProfile)
	if err != nil {
		return err
	}

	err = task.modListCreator.SaveAllToList(nil)
	if err != nil {
		return err
	}

	mainConfig.UpdateSelectedProfile(vanillaProfile.Name())

	err = task.mainConfigCreator.Write(mainConfig)
	if err != nil {
		return err
	}

	return nil
}
