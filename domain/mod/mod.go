package mod

import (
	"fmt"
	"strings"
)

type ModDto struct {
	Name         string
	Version      string
	Author       string
	Description  string
	Files        []FileDto
	Dependencies []string
}

type Mod struct {
	name         string
	version      string
	author       string
	description  string
	files        []File
	dependencies []Slug
}

func NewMod(dto ModDto) (Mod, error) {
	if strings.TrimSpace(dto.Name) == "" {
		return Mod{}, fmt.Errorf("mod must have name")
	}

	if strings.TrimSpace(dto.Version) == "" {
		return Mod{}, fmt.Errorf("mod must have version")
	}

	if strings.TrimSpace(dto.Author) == "" {
		return Mod{}, fmt.Errorf("mod must have author")
	}

	if strings.TrimSpace(dto.Name) == "" {
		return Mod{}, fmt.Errorf("mod must have description")
	}

	if len(dto.Files) == 0 {
		return Mod{}, fmt.Errorf("mod must have files")
	}

	files := make([]File, len(dto.Files))
	for i, fileDto := range dto.Files {
		file, err := NewFile(fileDto)
		if err != nil {
			return Mod{}, err
		}

		files[i] = file
	}

	var deps []Slug
	for _, dep := range dto.Dependencies {
		depSlug, err := NewSlugFromString(dep)
		if err != nil {
			return Mod{}, err
		}

		deps = append(deps, depSlug)
	}

	return Mod{
		name:         dto.Name,
		version:      dto.Version,
		author:       dto.Author,
		description:  dto.Description,
		files:        files,
		dependencies: deps,
	}, nil
}

func ReformMod(dto ModDto) Mod {
	files := make([]File, len(dto.Files))
	for i, dto := range dto.Files {
		files[i] = ReformFile(dto)
	}

	deps := make([]Slug, len(dto.Dependencies))
	for i, dep := range dto.Dependencies {
		deps[i] = ReformSlugFromString(dep)
	}

	return Mod{
		name:         dto.Name,
		version:      dto.Version,
		author:       dto.Author,
		description:  dto.Description,
		files:        files,
		dependencies: deps,
	}
}

func (mod Mod) Name() string {
	return mod.name
}

func (mod Mod) Version() string {
	return mod.version
}

func (mod Mod) Author() string {
	return mod.author
}

func (mod Mod) Description() string {
	return mod.description
}

func (mod Mod) Files() []File {
	return mod.files
}

func (mod Mod) Dependencies() []Slug {
	return mod.dependencies
}

func (mod Mod) Slug() Slug {
	return ReformSlug(mod.name, mod.author, mod.version)
}

func (mod Mod) Dto() ModDto {
	fileDtos := make([]FileDto, len(mod.files))
	for i, file := range mod.files {
		fileDtos[i] = file.Dto()
	}

	depStrings := make([]string, len(mod.dependencies))
	for i, dep := range mod.dependencies {
		depStrings[i] = dep.String()
	}

	return ModDto{
		Name:         mod.name,
		Version:      mod.version,
		Author:       mod.author,
		Description:  mod.description,
		Files:        fileDtos,
		Dependencies: depStrings,
	}
}
