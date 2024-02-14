package profile

import (
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
)

type profileDao interface {
	ReadByName(name string) (profile.ProfileDto, error)
}

type modDao interface {
	GetAllBySlugs(slugs []string) ([]mod.ModDto, error)
}

type ProfileTranslator struct {
	profileDao profileDao
	modDao     modDao
}

func NewProfileTranslator(
	profileDao profileDao,
	modDao modDao,
) ProfileTranslator {
	return ProfileTranslator{
		profileDao: profileDao,
		modDao:     modDao,
	}
}

func (translator ProfileTranslator) GetByName(name string) (profile.Profile, error) {
	profileDto, err := translator.profileDao.ReadByName(name)
	if err != nil {
		return profile.Profile{}, err
	}

	profileDto.Mods, err = translator.modDao.GetAllBySlugs(profileDto.ModNames)
	if err != nil {
		return profile.Profile{}, err
	}

	return profile.ReformProfile(profileDto), nil
}
