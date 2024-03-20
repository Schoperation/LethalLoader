package mod

import "schoperation/lethalloader/domain/mod"

type gameFilesDao interface {
	DeleteFilesByMod(mod mod.ModDto, gameFilesPath string) error
	AddFilesByMod(mod mod.ModDto, gameFilesPath string) error
}

type GameFilesTranslator struct {
	gameFilesDao gameFilesDao
}

func NewGameFilesTranslator(
	gameFilesDao gameFilesDao,
) GameFilesTranslator {
	return GameFilesTranslator{
		gameFilesDao: gameFilesDao,
	}
}

func (translator GameFilesTranslator) DeleteMod(mod mod.Mod, gameFilesPath string) error {
	err := translator.gameFilesDao.DeleteFilesByMod(mod.Dto(), gameFilesPath)
	if err != nil {
		return err
	}

	return nil
}

func (translator GameFilesTranslator) AddMod(mod mod.Mod, gameFilesPath string) error {
	err := translator.gameFilesDao.AddFilesByMod(mod.Dto(), gameFilesPath)
	if err != nil {
		return err
	}

	return nil
}
