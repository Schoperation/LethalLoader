package rest

import (
	"crypto/sha256"
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
	err := os.Mkdir("modcache", 0755)
	if err != nil {
		return mod.FileDto{}, err
	}

	file, err := os.Create("modcache/" + fileName)
	if err != nil {
		return mod.FileDto{}, err
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		return mod.FileDto{}, err
	}
	defer resp.Body.Close()

	hasher := sha256.New()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return mod.FileDto{}, err
	}

	_, err = io.Copy(hasher, resp.Body)
	if err != nil {
		return mod.FileDto{}, err
	}

	absFilePath, err := filepath.Abs("modcache/" + fileName)
	if err != nil {
		return mod.FileDto{}, err
	}

	return mod.FileDto{
		Name:      fileName,
		Path:      absFilePath,
		Sha256Sum: string(hasher.Sum(nil)),
	}, nil
}
