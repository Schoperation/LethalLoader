package file

import (
	"schoperation/lethalloader/domain/profile"
	"strings"
)

type ExportedProfileDao struct{}

func NewExportedDao() ExportedProfileDao {
	return ExportedProfileDao{}
}

const profilesDirectory = "profiles"

type exportedProfileModel struct {
	Name string   `json:"name"`
	Mods []string `json:"mods"`
}

func (model exportedProfileModel) dto() profile.ExportedProfileDto {
	return profile.ExportedProfileDto{
		Name:     model.Name,
		ModSlugs: model.Mods,
	}
}

func (model exportedProfileModel) key() string {
	lowered := strings.ToLower(model.Name)
	return strings.ReplaceAll(lowered, " ", "_")
}

func (dao ExportedProfileDao) GetAll() ([]profile.ExportedProfileDto, error) {
	models, err := readAllInDir[exportedProfileModel](profilesDirectory)
	if err != nil {
		return nil, err
	}

	dtos := make([]profile.ExportedProfileDto, len(models))
	i := 0
	for _, model := range models {
		dtos[i] = model.dto()
		i++
	}

	return dtos, nil
}

func (dao ExportedProfileDao) Save(dto profile.ExportedProfileDto) (string, error) {
	model := exportedProfileModel{
		Name: dto.Name,
		Mods: dto.ModSlugs,
	}

	err := write("profiles/pf_"+model.key()+".json", map[string]exportedProfileModel{model.key(): model})
	if err != nil {
		return "", err
	}

	return "profiles/pf_" + model.key() + ".json", nil
}
