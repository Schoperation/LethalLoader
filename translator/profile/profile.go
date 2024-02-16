package profile

import (
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
)

type profileDao interface {
	ReadByName(name string) (profile.ProfileDto, error)
	Write(dto profile.ProfileDto) error
}

type modListDao interface {
	GetAllBySlugs(slugs []string) ([]mod.ModDto, error)
}

type ProfileTranslator struct {
	profileDao profileDao
	modListDao modListDao
}

func NewProfileTranslator(
	profileDao profileDao,
	modListDao modListDao,
) ProfileTranslator {
	return ProfileTranslator{
		profileDao: profileDao,
		modListDao: modListDao,
	}
}

func (translator ProfileTranslator) GetByName(name string) (profile.Profile, error) {
	profileDto, err := translator.profileDao.ReadByName(name)
	if err != nil {
		return profile.Profile{}, err
	}

	profileDto.Mods, err = translator.modListDao.GetAllBySlugs(profileDto.ModSlugs)
	if err != nil {
		return profile.Profile{}, err
	}

	return profile.ReformProfile(profileDto), nil
}

func (translator ProfileTranslator) Save(pf profile.Profile) error {
	err := translator.profileDao.Write(pf.Dto())
	if err != nil {
		return err
	}

	return nil
}
