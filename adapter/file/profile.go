package file

import (
	"encoding/json"
	"os"
	"schoperation/lethalloader/domain/profile"
)

type ProfileDao struct{}

func NewProfileDao() ProfileDao {
	return ProfileDao{}
}

type profileModel struct {
	Name string   `json:"name"`
	Mods []string `json:"mods"`
}

func (dao ProfileDao) ReadAll() ([]profile.ProfileDto, error) {
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
		dtos[i] = profile.ProfileDto{
			Name:     model.Name,
			ModSlugs: model.Mods,
		}
		i++
	}

	return dtos, nil
}

func (dao ProfileDao) ReadByName(name string) (profile.ProfileDto, error) {
	file, err := os.ReadFile("profiles.json")
	if err != nil {
		return profile.ProfileDto{}, err
	}

	var model profileModel
	err = json.Unmarshal(file, &model)
	if err != nil {
		return profile.ProfileDto{}, err
	}

	return profile.ProfileDto{
		Name:     model.Name,
		ModSlugs: model.Mods,
	}, nil
}

func (dao ProfileDao) Write(dto profile.ProfileDto) error {
	model := profileModel{
		Name: dto.Name,
		Mods: dto.ModSlugs,
	}

	file, err := os.Create("profiles.json")
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(model, "", "    ")
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
