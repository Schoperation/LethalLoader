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

func (dao MainConfigDao) Read() (config.MainConfigDto, error) {
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
		GameFilePath:    model.GameFilePath,
		SelectedProfile: model.SelectedProfile,
	}, nil
}

func (dao MainConfigDao) Write(dto config.MainConfigDto) error {
	model := mainConfigModel{
		GameFilePath:    dto.GameFilePath,
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
