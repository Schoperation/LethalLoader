package mod

import (
	"schoperation/lethalloader/domain/mod"
)

type modDownloader interface {
	Download(url string, fileName string) (mod.FileDto, error)
}

type ModTranslator struct {
	modDownloader modDownloader
}

func NewModTranslator(modDownloader modDownloader) ModTranslator {
	return ModTranslator{
		modDownloader: modDownloader,
	}
}

const bepinex = "https://github.com/BepInEx/BepInEx/releases/download/v5.4.22/BepInEx_x64_5.4.22.0.zip"

func (translator ModTranslator) GetBepinEx() {

}
