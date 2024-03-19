package file

import (
	"errors"
	"os"
	"schoperation/lethalloader/domain/config"

	"gopkg.in/yaml.v3"
)

type MainConfigDao struct {
}

func NewMainConfigDao() MainConfigDao {
	return MainConfigDao{}
}

type mainConfigModel struct {
	GameFilePath    string `yaml:"gameFilePath"`
	SelectedProfile string `yaml:"selectedProfile"`
}

func (dao MainConfigDao) Get() (config.MainConfigDto, error) {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return config.MainConfigDto{}, nil
		}

		return config.MainConfigDto{}, err
	}

	var model mainConfigModel
	err = yaml.Unmarshal(file, &model)
	if err != nil {
		return config.MainConfigDto{}, err
	}

	return config.MainConfigDto{
		GameFilesPath:   model.GameFilePath,
		SelectedProfile: model.SelectedProfile,
	}, nil
}

func (dao MainConfigDao) Save(dto config.MainConfigDto) error {
	model := mainConfigModel{
		GameFilePath:    dto.GameFilesPath,
		SelectedProfile: dto.SelectedProfile,
	}

	bytes, err := yaml.Marshal(model)
	if err != nil {
		return err
	}

	err = os.WriteFile("config.yaml", bytes, 0660)
	if err != nil {
		return err
	}

	return nil
}
