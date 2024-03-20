package profile

import (
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
	"slices"
)

type profileDao interface {
	GetAll() ([]profile.ProfileDto, error)
	Save(dto profile.ProfileDto) error
	Delete(dto profile.ProfileDto) error
}

type modListDao interface {
	GetAllBySlugs(slugs []string) ([]mod.ModDto, error)
}

type gameFilesDao interface {
	DeleteFilesByProfile(pf profile.ProfileDto, gameFilesPath string) error
	AddFilesByProfile(pf profile.ProfileDto, gameFilesPath string) error
}

type ProfileTranslator struct {
	profileDao   profileDao
	modListDao   modListDao
	gameFilesDao gameFilesDao
}

func NewProfileTranslator(
	profileDao profileDao,
	modListDao modListDao,
	gameFilesDao gameFilesDao,
) ProfileTranslator {
	return ProfileTranslator{
		profileDao:   profileDao,
		modListDao:   modListDao,
		gameFilesDao: gameFilesDao,
	}
}

func (translator ProfileTranslator) GetAll() ([]profile.Profile, error) {
	profileDtos, err := translator.profileDao.GetAll()
	if err != nil {
		return nil, err
	}

	profiles := make([]profile.Profile, len(profileDtos))
	for i, dto := range profileDtos {
		dto.Mods, err = translator.modListDao.GetAllBySlugs(dto.ModSlugs)
		if err != nil {
			return nil, err
		}

		profiles[i] = profile.ReformProfile(dto)
	}

	slices.SortFunc(profiles, func(a, b profile.Profile) int {
		if a.Name() < b.Name() {
			return -1
		}

		if a.Name() > b.Name() {
			return 1
		}

		return 0
	})

	return profiles, nil
}

func (translator ProfileTranslator) Save(pf profile.Profile) error {
	err := translator.profileDao.Save(pf.Dto())
	if err != nil {
		return err
	}

	return nil
}

func (translator ProfileTranslator) Delete(pf profile.Profile) error {
	err := translator.profileDao.Delete(pf.Dto())
	if err != nil {
		return err
	}

	return nil
}

func (translator ProfileTranslator) Switch(oldPf profile.Profile, newPf profile.Profile, gameFilesPath string) error {
	err := translator.gameFilesDao.DeleteFilesByProfile(oldPf.Dto(), gameFilesPath)
	if err != nil {
		return err
	}

	err = translator.gameFilesDao.AddFilesByProfile(newPf.Dto(), gameFilesPath)
	if err != nil {
		return err
	}

	return nil
}
