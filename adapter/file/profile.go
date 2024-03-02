package file

import (
	"schoperation/lethalloader/domain/profile"
	"strings"
)

type ProfileDao struct{}

func NewProfileDao() ProfileDao {
	return ProfileDao{}
}

const profilesFileName = "profiles.json"

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
	models, err := read[profileModel](profilesFileName)
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

	models, err := read[profileModel](profilesFileName)
	if err != nil {
		return err
	}

	models[model.key()] = model

	err = write(profilesFileName, models)
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

	models, err := read[profileModel](profilesFileName)
	if err != nil {
		return err
	}

	delete(models, modelToDelete.key())

	err = write(profilesFileName, models)
	if err != nil {
		return err
	}

	return nil
}
