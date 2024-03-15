package file

import (
	"fmt"
	"schoperation/lethalloader/domain/mod"
	"slices"
	"strings"
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
		Dependencies: model.Dependencies,
		Files:        fileDtos,
	}
}

type fileModel struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Sha256Sum string `json:"sha256sum"`
}

func (model fileModel) dto() mod.FileDto {
	return mod.FileDto{
		Name:      model.Name,
		Path:      model.Path,
		Sha256Sum: model.Sha256Sum,
	}
}

func (dao ModListDao) slug(name, author, version string) string {
	return author + "-" + name + "-" + version
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

	dtos := make([]mod.ModDto, len(models))
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

func (dao ModListDao) GetByNameAuthorVersion(name, author, version string) (mod.ModDto, error) {
	models, err := read[modModel](modListFileName)
	if err != nil {
		return mod.ModDto{}, err
	}

	slug := dao.slug(name, author, version)
	for slugKey, model := range models {
		if slugKey == slug {
			return model.dto(), nil
		}
	}

	return mod.ModDto{}, fmt.Errorf("mod not found")
}

func (dao ModListDao) GetBySearchTerm(term string) ([]mod.ModDto, error) {
	term = strings.ToLower(term)
	term = strings.ReplaceAll(term, " ", "")

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
	}

	return foundDtos, nil
}

func (dao ModListDao) Save(dto mod.ModDto) error {
	fileModels := make([]fileModel, len(dto.Files))
	for j, fileDto := range dto.Files {
		fileModels[j] = fileModel{
			Name:      fileDto.Name,
			Path:      fileDto.Path,
			Sha256Sum: fileDto.Sha256Sum,
		}
	}

	newModModel := modModel{
		Name:         dto.Name,
		Version:      dto.Version,
		Author:       dto.Author,
		Description:  dto.Description,
		Dependencies: dto.Dependencies,
		Files:        fileModels,
	}

	models, err := read[modModel](modListFileName)
	if err != nil {
		return err
	}

	slug := dao.slug(dto.Name, dto.Author, dto.Version)
	models[slug] = newModModel

	err = write(modListFileName, models)
	if err != nil {
		return err
	}

	return nil
}
