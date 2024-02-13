package rest

import (
	"net/http"
	"os"
	"schoperation/lethalloader/domain/mod"
)

type BepinExDownloader struct{}

func NewBepinExDownloader() BepinExDownloader {
	return BepinExDownloader{}
}

func (dldr BepinExDownloader) Get() (mod.FileDto, error) {
	file, err := os.Create("/modcache/bepinex.zip")
	if err != nil {
		return mod.FileDto{}, err
	}

	resp, err := http.Get("https://github.com/BepInEx/BepInEx/releases/download/v5.4.22/BepInEx_x64_5.4.22.0.zip")
	if err != nil {
		return mod.FileDto{}, err
	}

}
