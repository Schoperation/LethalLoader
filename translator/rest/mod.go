package rest

import "schoperation/lethalloader/domain/mod"

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

func (translator ModTranslator) Get(name string) (mod.Mod, error) {

}
