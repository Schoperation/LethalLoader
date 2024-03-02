package file

import (
	"encoding/json"
	"fmt"
	"io"
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

func (dao ModListDao) slug(name, author, version string) string {
	return author + "-" + name + "-" + version
}

func (dao ModListDao) getModModels() (map[string]modModel, error) {
	file, err := os.Create("mods.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if len(bytes) == 0 {
		return make(map[string]modModel), nil
	}

	models := make(map[string]modModel)
	err = json.Unmarshal(bytes, &models)
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (dao ModListDao) GetAll() ([]mod.ModDto, error) {
	models, err := dao.getModModels()
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
	models, err := dao.getModModels()
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

func (dao ModListDao) GetByNameAuthorVersion(name, author, version string) (mod.ModDto, error) {
	models, err := dao.getModModels()
	if err != nil {
		return mod.ModDto{}, err
	}

	slug := dao.slug(name, author, version)
	for slugKey, model := range models {
		if slugKey == slug {
			return model.Dto(), nil
		}
	}

	return mod.ModDto{}, fmt.Errorf("mod not found")
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

	models, err := dao.getModModels()
	if err != nil {
		return err
	}

	slug := dao.slug(dto.Name, dto.Author, dto.Version)
	models[slug] = newModModel

	bytes, err := json.MarshalIndent(models, "", "    ")
	if err != nil {
		return err
	}

	file, err := os.Create("mods.json")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
