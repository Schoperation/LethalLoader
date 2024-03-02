package file

import (
	"encoding/json"
	"io"
	"os"
	"schoperation/lethalloader/domain/profile"
	"strings"
)

type ProfileDao struct{}

func NewProfileDao() ProfileDao {
	return ProfileDao{}
}

type profileModel struct {
	Name string   `json:"name"`
	Mods []string `json:"mods"`
}

func (model profileModel) dto() profile.ProfileDto {
	return profile.ProfileDto{
		Name:     model.Name,
		ModSlugs: model.Mods,
	}
}

func (model profileModel) key() string {
	lowered := strings.ToLower(model.Name)
	return strings.ReplaceAll(lowered, " ", "_")
}

func (dao ProfileDao) GetAll() ([]profile.ProfileDto, error) {
	file, err := os.ReadFile("profiles.json")
	if err != nil {
		return nil, err
	}

	models := make(map[string]profileModel)
	err = json.Unmarshal(file, &models)
	if err != nil {
		return nil, err
	}

	dtos := make([]profile.ProfileDto, len(models))
	i := 0
	for _, model := range models {
		dtos[i] = model.dto()
		i++
	}

	return dtos, nil
}

func (dao ProfileDao) Save(dto profile.ProfileDto) error {
	model := profileModel{
		Name: dto.Name,
		Mods: dto.ModSlugs,
	}

	file, err := os.Create("profiles.json")
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	models := make(map[string]profileModel)
	err = json.Unmarshal(bytes, &models)
	if err != nil {
		return err
	}

	models[model.key()] = model

	bytes, err = json.MarshalIndent(models, "", "    ")
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func (dao ProfileDao) Delete(dto profile.ProfileDto) error {
	modelToDelete := profileModel{
		Name: dto.Name,
		Mods: dto.ModSlugs,
	}

	file, err := os.Create("profiles.json")
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	models := make(map[string]profileModel)
	err = json.Unmarshal(bytes, &models)
	if err != nil {
		return err
	}

	delete(models, modelToDelete.key())

	bytes, err = json.MarshalIndent(models, "", "    ")
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
