package mod

import (
	"fmt"
	"strings"
)

type FileDto struct {
	Name      string
	Path      string
	Sha256Sum string
}

type File struct {
	name      string
	path      string
	sha256sum string
}

func NewFile(dto FileDto) (File, error) {
	if strings.TrimSpace(dto.Name) == "" {
		return File{}, fmt.Errorf("file must have name")
	}

	if strings.TrimSpace(dto.Path) == "" {
		return File{}, fmt.Errorf("file must have path")
	}

	dto.Path = strings.Replace(dto.Path, dto.Name, "", 1)

	if strings.TrimSpace(dto.Sha256Sum) == "" {
		return File{}, fmt.Errorf("file must have sum")
	}

	return File{
		name:      dto.Name,
		path:      dto.Path,
		sha256sum: dto.Sha256Sum,
	}, nil
}

func ReformFile(dto FileDto) File {
	return File{
		name:      dto.Name,
		path:      dto.Path,
		sha256sum: dto.Sha256Sum,
	}
}

func (file File) Name() string {
	return file.name
}

func (file File) Path() string {
	return file.path
}

func (file File) Sha256Sum() string {
	return file.sha256sum
}
