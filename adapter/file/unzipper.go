package file

import (
	"archive/zip"
	"schoperation/lethalloader/domain/mod"
)

type FileUnzipper struct{}

func NewFileUnzipper() FileUnzipper {
	return FileUnzipper{}
}

func (fuz FileUnzipper) Unzip(zippedDto mod.FileDto) ([]mod.FileDto, error) {
	reader, err := zip.OpenReader(zippedDto.Path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return nil, nil
}
