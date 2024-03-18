package file

import (
	"archive/zip"
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

func (fuzpr FileUnzipper) Unzip(zippedDto mod.FileDto) ([]mod.FileDto, error) {
	reader, err := zip.OpenReader(zippedDto.Path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	defer os.Remove(zippedDto.Path)

	err = os.MkdirAll("modcache"+string(os.PathSeparator)+zippedDto.Name, 0755)
	if err != nil {
		return nil, err
	}

	fileDtos := []mod.FileDto{}
	for _, f := range reader.File {
		fileDto, err := fuzpr.extractFile(f, "modcache"+string(os.PathSeparator)+zippedDto.Name)
		if err != nil {
			return nil, err
		}

		if fileDto.Name == "skip" {
			continue
		}

		fileDtos = append(fileDtos, fileDto)
	}

	return fileDtos, nil
}

// TODO take care of edge cases
func (fuz FileUnzipper) extractFile(file *zip.File, unzippedFolder string) (mod.FileDto, error) {
	rc, err := file.Open()
	if err != nil {
		return mod.FileDto{}, err
	}
	defer rc.Close()

	// author-modname-version
	names := strings.Split(unzippedFolder, "-")
	modName := names[1]

	path := filepath.Join(unzippedFolder, file.Name)

	// Remove extraneous modname folders.
	if file.Name == modName+string(os.PathSeparator) {
		return mod.FileDto{Name: "skip"}, nil
	}

	// Remove files not part of mod (readme, manifest, icon, etc.). They are at the base of the zip file.
	if strings.TrimSuffix(path, string(os.PathSeparator)+file.FileInfo().Name()) == unzippedFolder && !file.FileInfo().IsDir() {
		return mod.FileDto{Name: "skip"}, nil
	}

	// Check for ZipSlip.
	if !strings.HasPrefix(path, filepath.Clean(unzippedFolder)+string(os.PathSeparator)) {
		return mod.FileDto{}, fmt.Errorf("illegal file path: %s", path)
	}

	path = strings.Replace(path, modName+string(os.PathSeparator), "", 1)

	if file.FileInfo().IsDir() {
		os.MkdirAll(path, file.Mode())
		return mod.FileDto{Name: "skip"}, nil
	}

	os.Mkdir(filepath.Dir(path), file.Mode())

	f, err := os.Create(path)
	if err != nil {
		return mod.FileDto{}, err
	}
	defer f.Close()

	_, err = io.Copy(f, rc)
	if err != nil {
		return mod.FileDto{}, err
	}

	stats, err := f.Stat()
	if err != nil {
		return mod.FileDto{}, err
	}

	return mod.FileDto{
		Name: stats.Name(),
		Path: strings.TrimPrefix(path, unzippedFolder+string(os.PathSeparator)),
	}, nil
}
