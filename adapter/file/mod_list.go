package file

import (
	"encoding/json"
	"fmt"
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

func (model modModel) Dto() mod.ModDto {
	fileDtos := make([]mod.FileDto, len(model.Files))
	for i, fileModel := range model.Files {
		fileDtos[i] = fileModel.Dto()
	}

	return mod.ModDto{
		Name:         model.Name,
		Version:      model.Version,
		Author:       model.Author,
		Description:  model.Description,
		Dependencies: model.Dependencies,
		Files:        fileDtos,
	}
}

type fileModel struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Sha256Sum string `json:"sha256sum"`
}

func (model fileModel) Dto() mod.FileDto {
	return mod.FileDto{
		Name:      model.Name,
		Path:      model.Path,
		Sha256Sum: model.Sha256Sum,
	}
}

func (dao ModListDao) GetBySlug(slug string) (mod.ModDto, error) {
	file, err := os.ReadFile("mods.json")
	if err != nil {
		return mod.ModDto{}, err
	}

	models := make(map[string]modModel)
	err = json.Unmarshal(file, &models)
	if err != nil {
		return mod.ModDto{}, err
	}

	for slugKey, model := range models {
		if slugKey == slug {
			return model.Dto(), nil
		}
	}

	return mod.ModDto{}, fmt.Errorf("mod not found")
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

	dtos := make([]mod.ModDto, len(models))
	i := 0
	for _, model := range models {
		dtos[i] = model.Dto()
		i++
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

	dtos := make([]mod.ModDto, len(models))
	i := 0
	for slug, model := range models {
		if !slices.Contains(slugs, slug) {
			continue
		}

		dtos[i] = model.Dto()
		i++
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
	defer file.Close()

	if len(modModels) == 0 {
		_, err = file.WriteString("{}")
		if err != nil {
			return err
		}

		return nil
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
