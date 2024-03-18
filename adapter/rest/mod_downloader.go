package rest

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"schoperation/lethalloader/domain/mod"
)

type ModDownloader struct{}

func NewModDownloader() ModDownloader {
	return ModDownloader{}
}

func (dldr ModDownloader) Download(url string, fileName string) (mod.FileDto, error) {
	err := os.MkdirAll("zips", 0755)
	if err != nil {
		return mod.FileDto{}, err
	}

	file, err := os.Create("zips" + string(os.PathSeparator) + fileName)
	if err != nil {
		return mod.FileDto{}, err
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		return mod.FileDto{}, err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return mod.FileDto{}, err
	}

	absFilePath, err := filepath.Abs("zips" + string(os.PathSeparator) + fileName)
	if err != nil {
		return mod.FileDto{}, err
	}

	return mod.FileDto{
		Name: fileName,
		Path: absFilePath,
	}, nil
}
