package profile

import (
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
)

type profileDao interface {
	ReadAll() ([]profile.ProfileDto, error)
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

func (translator ProfileTranslator) GetAll(name string) ([]profile.Profile, error) {
	profileDtos, err := translator.profileDao.ReadAll()
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

	return profiles, nil
}

func (translator ProfileTranslator) Save(pf profile.Profile) error {
	err := translator.profileDao.Write(pf.Dto())
	if err != nil {
		return err
	}

	return nil
}
