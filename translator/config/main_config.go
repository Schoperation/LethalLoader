package config

import "schoperation/lethalloader/domain/config"

type mainConfigDao interface {
	Read() (config.MainConfigDto, error)
	Write(dto config.MainConfigDto) error
}

type MainConfigTranslator struct {
	mainConfigDao mainConfigDao
}

func NewMainConfigTranslator(mainConfigDao mainConfigDao) MainConfigTranslator {
	return MainConfigTranslator{
		mainConfigDao: mainConfigDao,
	}
}

func (translator MainConfigTranslator) Read() (config.MainConfig, error) {
	mainConfig, err := translator.mainConfigDao.Read()
	if err != nil {
		return config.MainConfig{}, err
	}

	return config.ReformMainConfig(mainConfig), nil
}

func (translator MainConfigTranslator) Write(mainConfig config.MainConfig) error {
	err := translator.mainConfigDao.Write(mainConfig.Dto())
	if err != nil {
		return err
	}

	return nil
}
