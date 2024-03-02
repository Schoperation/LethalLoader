package config

import "schoperation/lethalloader/domain/config"

type mainConfigDao interface {
	Get() (config.MainConfigDto, error)
	Save(dto config.MainConfigDto) error
}

type MainConfigTranslator struct {
	mainConfigDao mainConfigDao
}

func NewMainConfigTranslator(mainConfigDao mainConfigDao) MainConfigTranslator {
	return MainConfigTranslator{
		mainConfigDao: mainConfigDao,
	}
}

func (translator MainConfigTranslator) Get() (config.MainConfig, error) {
	mainConfig, err := translator.mainConfigDao.Get()
	if err != nil {
		return config.MainConfig{}, err
	}

	return config.ReformMainConfig(mainConfig), nil
}

func (translator MainConfigTranslator) Save(mainConfig config.MainConfig) error {
	err := translator.mainConfigDao.Save(mainConfig.Dto())
	if err != nil {
		return err
	}

	return nil
}
