package profile

import (
	"schoperation/lethalloader/domain/profile"
	"slices"
	"strings"
)

type exportedProfileDao interface {
	GetAll() ([]profile.ExportedProfileDto, error)
	Save(dto profile.ExportedProfileDto) (string, error)
}

type ExportedProfileTranslator struct {
	exportedProfileDao exportedProfileDao
}

func NewExportedProfileTranslator(
	exportedProfileDao exportedProfileDao,
) ExportedProfileTranslator {
	return ExportedProfileTranslator{
		exportedProfileDao: exportedProfileDao,
	}
}

func (translator ExportedProfileTranslator) GetAll() ([]profile.ExportedProfile, error) {
	dtos, err := translator.exportedProfileDao.GetAll()
	if err != nil {
		return nil, err
	}

	profiles := make([]profile.ExportedProfile, len(dtos))
	for i, dto := range dtos {
		profiles[i] = profile.ReformExportedProfile(dto)
	}

	slices.SortFunc(profiles, func(a, b profile.ExportedProfile) int {
		an := strings.ToLower(a.Name())
		bn := strings.ToLower(b.Name())

		if an < bn {
			return -1
		}

		if an > bn {
			return 1
		}

		return 0
	})

	return profiles, nil
}

func (translator ExportedProfileTranslator) Export(pf profile.Profile) (string, error) {
	fileName, err := translator.exportedProfileDao.Save(pf.ExportedDto())
	if err != nil {
		return "", err
	}

	return fileName, nil
}
