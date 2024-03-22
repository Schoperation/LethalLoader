package mod

import (
	"schoperation/lethalloader/domain/config"
	"schoperation/lethalloader/domain/mod"
)

type gameFilesDao interface {
	DeleteFilesByMod(mod mod.ModDto, gameFilesPath string) error
	AddFilesByMod(mod mod.ModDto, gameFilesPath string) error
}

type gameFilesPathConfigDao interface {
	Get() (config.MainConfigDto, error)
}

type GameFilesTranslator struct {
	gameFilesDao  gameFilesDao
	mainConfigDao gameFilesPathConfigDao
}

func NewGameFilesTranslator(
	gameFilesDao gameFilesDao,
	mainConfigDao gameFilesPathConfigDao,
) GameFilesTranslator {
	return GameFilesTranslator{
		gameFilesDao:  gameFilesDao,
		mainConfigDao: mainConfigDao,
	}
}

func (translator GameFilesTranslator) DeleteMod(mod mod.Mod, pfName string) error {
	mainConfig, err := translator.mainConfigDao.Get()
	if err != nil {
		return err
	}

	if mainConfig.SelectedProfile != pfName {
		return nil
	}

	err = translator.gameFilesDao.DeleteFilesByMod(mod.Dto(), mainConfig.GameFilesPath)
	if err != nil {
		return err
	}

	return nil
}

func (translator GameFilesTranslator) AddMod(mod mod.Mod, pfName string) error {
	mainConfig, err := translator.mainConfigDao.Get()
	if err != nil {
		return err
	}

	if mainConfig.SelectedProfile != pfName {
		return nil
	}

	err = translator.gameFilesDao.AddFilesByMod(mod.Dto(), mainConfig.GameFilesPath)
	if err != nil {
		return err
	}

	return nil
}
