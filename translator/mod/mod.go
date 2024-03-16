package mod

import (
	"schoperation/lethalloader/domain/mod"
	"strings"
)

type modDownloader interface {
	Download(url string, fileName string) (mod.FileDto, error)
}

type modUnzipper interface {
	Unzip(zippedDto mod.FileDto) ([]mod.FileDto, error)
}

type modListDao interface {
	GetByNameAuthorVersion(name, author, version string) (mod.ModDto, error)
	GetAllBySearchTerm(term string) ([]mod.ModDto, error)
	GetAll() ([]mod.ModDto, error)
	Save(dto mod.ModDto) error
}

type ModTranslator struct {
	modDownloader modDownloader
	modUnzipper   modUnzipper
	modListDao    modListDao
}

func NewModTranslator(
	modDownloader modDownloader,
	modUnzipper modUnzipper,
	modListDao modListDao,
) ModTranslator {
	return ModTranslator{
		modDownloader: modDownloader,
		modUnzipper:   modUnzipper,
		modListDao:    modListDao,
	}
}

func (translator ModTranslator) GetByModListing(listing mod.Listing) (mod.Mod, error) {
	cachedMod, err := translator.modListDao.GetByNameAuthorVersion(listing.Name(), listing.Author(), listing.Version())
	if err != nil && err.Error() != "mod not found" {
		return mod.Mod{}, err
	}

	if err == nil {
		return mod.ReformMod(cachedMod), nil
	}

	fileName := listing.Author() + "-" + listing.Name() + "-" + listing.Version()
	zipFile, err := translator.modDownloader.Download(listing.DownloadUrl(), fileName)
	if err != nil {
		return mod.Mod{}, err
	}

	fileDtos, err := translator.modUnzipper.Unzip(zipFile)
	if err != nil {
		return mod.Mod{}, err
	}

	newModDto := mod.ModDto{
		Name:         listing.Name(),
		Version:      listing.Version(),
		Author:       listing.Author(),
		Description:  listing.Description(),
		Files:        fileDtos,
		Dependencies: listing.Dependencies(),
	}

	newMod, err := mod.NewMod(newModDto)
	if err != nil {
		return mod.Mod{}, err
	}

	err = translator.modListDao.Save(newModDto)
	if err != nil {
		return mod.Mod{}, err
	}

	return newMod, nil
}

func (translator ModTranslator) GetAllBySearchTerm(term string) ([]mod.Mod, error) {
	term = strings.ToLower(term)
	term = strings.ReplaceAll(term, " ", "")

	dtos, err := translator.modListDao.GetAllBySearchTerm(term)
	if err != nil {
		return nil, err
	}

	mods := make([]mod.Mod, len(dtos))
	for i, dto := range dtos {
		mods[i] = mod.ReformMod(dto)
	}

	return mods, nil
}

func (translator ModTranslator) GetAllFromList() ([]mod.Mod, error) {
	dtos, err := translator.modListDao.GetAll()
	if err != nil {
		return nil, err
	}

	mods := make([]mod.Mod, len(dtos))
	for i, dto := range dtos {
		mods[i] = mod.ReformMod(dto)
	}

	return mods, nil
}

func (translator ModTranslator) SaveToList(mod mod.Mod) error {
	err := translator.modListDao.Save(mod.Dto())
	if err != nil {
		return err
	}

	return nil
}
