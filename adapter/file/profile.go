package file

import (
	"encoding/json"
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

func (dao ProfileDao) ReadByName(name string) (profile.ProfileDto, error) {
	file, err := os.ReadFile(dao.fileNameFromProfileName(name))
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
		ModNames: model.Mods,
	}, nil
}

func (dao ProfileDao) Write(dto profile.ProfileDto) error {
	model := profileModel{
		Name: dto.Name,
		Mods: dto.ModNames,
	}

	newFile, err := os.Create(dao.fileNameFromProfileName(dto.Name))
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(model, "", "    ")
	if err != nil {
		return err
	}

	_, err = newFile.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func (dao ProfileDao) fileNameFromProfileName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	return name
}
