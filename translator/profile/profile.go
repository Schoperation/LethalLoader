package profile

import (
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
	"slices"
)

type profileDao interface {
	GetAll() ([]profile.ProfileDto, error)
	Save(dto profile.ProfileDto) error
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
