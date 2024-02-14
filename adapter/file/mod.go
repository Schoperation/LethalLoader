package file

import (
	"encoding/json"
	"os"
	"schoperation/lethalloader/domain/mod"
	"slices"
)

type ModDao struct {
}

func NewModDao() ModDao {
	return ModDao{}
}

type modModel struct {
	Name         string      `json:"name"`
	Version      string      `json:"version"`
	Author       string      `json:"author"`
	Description  string      `json:"description"`
	Dependencies []string    `json:"dependencies"`
	Files        []fileModel `json:"files"`
}

type fileModel struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Sha256Sum string `json:"sha256sum"`
}

func (dao ModDao) GetAllBySlugs(slugs []string) ([]mod.ModDto, error) {
	file, err := os.ReadFile("mods.json")
	if err != nil {
		return nil, err
	}

	models := make(map[string]modModel)
	err = json.Unmarshal(file, &models)
	if err != nil {
		return nil, err
	}

	dtos := []mod.ModDto{}
	for slug, model := range models {
		if !slices.Contains(slugs, slug) {
			continue
		}

		fileDtos := make([]mod.FileDto, len(model.Files))
		for i, fileModel := range model.Files {
			fileDtos[i] = mod.FileDto{
				Name:      fileModel.Name,
				Path:      fileModel.Path,
				Sha256Sum: fileModel.Sha256Sum,
			}
		}

		dtos = append(dtos, mod.ModDto{
			Name:         model.Name,
			Version:      model.Version,
			Author:       model.Author,
			Description:  model.Description,
			Dependencies: model.Dependencies,
			Files:        fileDtos,
		})
	}

	return dtos, nil
}
