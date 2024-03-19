package mod

import (
	"fmt"
	"strings"
)

type FileDto struct {
	Name string
	Path string
}

type File struct {
	name string
	path string
}

func NewFile(dto FileDto) (File, error) {
	if strings.TrimSpace(dto.Name) == "" {
		return File{}, fmt.Errorf("file must have name")
	}

	// if strings.TrimSpace(dto.Path) == "" {
	// 	return File{}, fmt.Errorf("file must have path")
	// }

	dto.Path = strings.Replace(dto.Path, dto.Name, "", 1)

	return File{
		name: dto.Name,
		path: dto.Path,
	}, nil
}

func ReformFile(dto FileDto) File {
	return File{
		name: dto.Name,
		path: dto.Path,
	}
}

func (file File) Name() string {
	return file.name
}

func (file File) Path() string {
	return file.path
}

func (file File) Dto() FileDto {
	return FileDto{
		Name: file.name,
		Path: file.path,
	}
}
