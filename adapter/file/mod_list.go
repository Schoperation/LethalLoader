package file

import (
	"encoding/json"
	"os"
	"schoperation/lethalloader/domain/mod"
	"slices"
)

type ModListDao struct {
}

func NewModListDao() ModListDao {
	return ModListDao{}
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

func (dao ModListDao) GetAll() ([]mod.ModDto, error) {
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
	for _, model := range models {
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

func (dao ModListDao) GetAllBySlugs(slugs []string) ([]mod.ModDto, error) {
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

func (dao ModListDao) SaveAll(dtos []mod.ModDto) error {
	modModels := make([]modModel, len(dtos))
	for i, dto := range dtos {
		fileModels := make([]fileModel, len(dto.Files))
		for j, fileDto := range dto.Files {
			fileModels[j] = fileModel{
				Name:      fileDto.Name,
				Path:      fileDto.Path,
				Sha256Sum: fileDto.Sha256Sum,
			}
		}

		modModels[i] = modModel{
			Name:         dto.Name,
			Version:      dto.Version,
			Author:       dto.Author,
			Description:  dto.Description,
			Dependencies: dto.Dependencies,
			Files:        fileModels,
		}
	}

	file, err := os.Create("mods.json")
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(modModels, "", "    ")
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
