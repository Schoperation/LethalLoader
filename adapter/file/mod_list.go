package file

import (
	"fmt"
	"schoperation/lethalloader/domain/mod"
	"slices"
	"strings"
	"time"
)

type ModListDao struct {
}

func NewModListDao() ModListDao {
	return ModListDao{}
}

const modListFileName = "mods.json"

type modModel struct {
	Name         string      `json:"name"`
	Version      string      `json:"version"`
	Author       string      `json:"author"`
	Description  string      `json:"description"`
	DateCreated  time.Time   `json:"date_created"`
	Dependencies []string    `json:"dependencies"`
	Files        []fileModel `json:"files"`
}

func (model modModel) dto() mod.ModDto {
	fileDtos := make([]mod.FileDto, len(model.Files))
	for i, fileModel := range model.Files {
		fileDtos[i] = fileModel.dto()
	}

	return mod.ModDto{
		Name:         model.Name,
		Version:      model.Version,
		Author:       model.Author,
		Description:  model.Description,
		DateCreated:  model.DateCreated,
		Dependencies: model.Dependencies,
		Files:        fileDtos,
	}
}

type fileModel struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (model fileModel) dto() mod.FileDto {
	return mod.FileDto{
		Name: model.Name,
		Path: model.Path,
	}
}

func (dao ModListDao) GetAll() ([]mod.ModDto, error) {
	models, err := read[modModel](modListFileName)
	if err != nil {
		return nil, err
	}

	dtos := make([]mod.ModDto, len(models))
	i := 0
	for _, model := range models {
		dtos[i] = model.dto()
		i++
	}

	return dtos, nil
}

func (dao ModListDao) GetAllBySlugs(slugs []string) ([]mod.ModDto, error) {
	if len(slugs) == 0 {
		return []mod.ModDto{}, nil
	}

	models, err := read[modModel](modListFileName)
	if err != nil {
		return nil, err
	}

	dtos := make([]mod.ModDto, len(slugs))
	i := 0
	for slug, model := range models {
		if !slices.Contains(slugs, slug) {
			continue
		}

		dtos[i] = model.dto()
		i++
	}

	return dtos, nil
}

func (dao ModListDao) GetBySlug(slug string) (mod.ModDto, error) {
	models, err := read[modModel](modListFileName)
	if err != nil {
		return mod.ModDto{}, err
	}

	for slugKey, model := range models {
		if slugKey == slug {
			return model.dto(), nil
		}
	}

	return mod.ModDto{}, fmt.Errorf("mod not found")
}

func (dao ModListDao) GetAllBySearchTerm(term string) ([]mod.ModDto, error) {
	models, err := read[modModel](modListFileName)
	if err != nil {
		return nil, err
	}

	var foundDtos []mod.ModDto
	for slug, model := range models {
		slug = strings.ToLower(slug)

		if strings.Contains(slug, term) {
			foundDtos = append(foundDtos, model.dto())
		}

		if len(foundDtos) == 10 {
			break
		}
	}

	return foundDtos, nil
}

func (dao ModListDao) Save(dto mod.ModDto, slug string) error {
	fileModels := make([]fileModel, len(dto.Files))
	for j, fileDto := range dto.Files {
		fileModels[j] = fileModel{
			Name: fileDto.Name,
			Path: fileDto.Path,
		}
	}

	newModModel := modModel{
		Name:         dto.Name,
		Version:      dto.Version,
		Author:       dto.Author,
		Description:  dto.Description,
		DateCreated:  dto.DateCreated,
		Dependencies: dto.Dependencies,
		Files:        fileModels,
	}

	models, err := read[modModel](modListFileName)
	if err != nil {
		return err
	}

	models[slug] = newModModel

	err = write(modListFileName, models)
	if err != nil {
		return err
	}

	return nil
}
