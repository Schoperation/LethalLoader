package mod

import (
	"schoperation/lethalloader/domain/mod"
)

type modDownloader interface {
	Download(url string, fileName string) (mod.FileDto, error)
}

type modListDao interface {
	GetAll() ([]mod.ModDto, error)
	SaveAll(dtos []mod.ModDto) error
}

type ModTranslator struct {
	modDownloader modDownloader
	modListDao    modListDao
}

func NewModTranslator(
	modDownloader modDownloader,
	modListDao modListDao,
) ModTranslator {
	return ModTranslator{
		modDownloader: modDownloader,
		modListDao:    modListDao,
	}
}

const bepinex = "https://github.com/BepInEx/BepInEx/releases/download/v5.4.22/BepInEx_x64_5.4.22.0.zip"

func (translator ModTranslator) GetBepinEx() {

}

func (translator ModTranslator) Download() {

}

func (translator ModTranslator) GetAllFromList() ([]mod.Mod, error) {
	dtos, err := translator.modListDao.GetAll()
	if err != nil {
		return nil, err
	}

	mods := make([]mod.Mod, len(dtos))
	for i, dto := range dtos {
		mods[i] = mod.ReformMod(dto)
	}

	return mods, nil
}

func (translator ModTranslator) SaveAllToList(mods []mod.Mod) error {
	dtos := make([]mod.ModDto, len(mods))
	for i, mod := range mods {
		dtos[i] = mod.Dto()
	}

	err := translator.modListDao.SaveAll(dtos)
	if err != nil {
		return err
	}

	return nil
}
