package file

import (
	"archive/zip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"schoperation/lethalloader/domain/mod"
	"strings"
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

	err = os.MkdirAll("modcache/"+zippedDto.Name, 0755)
	if err != nil {
		return nil, err
	}

	fileDtos := []mod.FileDto{}
	for _, f := range reader.File {
		fileDto, err := fuz.extractFile(f, "modcache/"+zippedDto.Name)
		if err != nil {
			return nil, err
		}

		if fileDto.Name == "directory" {
			continue
		}

		fileDtos = append(fileDtos, fileDto)
	}

	return fileDtos, nil
}

func (fuz FileUnzipper) extractFile(file *zip.File, unzippedFolder string) (mod.FileDto, error) {
	rc, err := file.Open()
	if err != nil {
		return mod.FileDto{}, err
	}
	defer rc.Close()

	path := filepath.Join(unzippedFolder, file.Name)

	// Check for ZipSlip
	if !strings.HasPrefix(path, filepath.Clean(unzippedFolder)+string(os.PathSeparator)) {
		return mod.FileDto{}, fmt.Errorf("illegal file path: %s", path)
	}

	if file.FileInfo().IsDir() {
		os.MkdirAll(path, file.Mode())
		return mod.FileDto{Name: "directory"}, nil
	}

	os.Mkdir(filepath.Dir(path), file.Mode())
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return mod.FileDto{}, err
	}
	defer f.Close()

	_, err = io.Copy(f, rc)
	if err != nil {
		return mod.FileDto{}, err
	}

	hasher := sha256.New()
	_, err = io.Copy(hasher, rc)
	if err != nil {
		return mod.FileDto{}, err
	}

	return mod.FileDto{
		Name:      f.Name(),
		Path:      path,
		Sha256Sum: string(hasher.Sum(nil)),
	}, nil
}
