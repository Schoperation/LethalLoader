package profile

import "schoperation/lethalloader/domain/mod"

type ExportedProfileDto struct {
	Name     string
	ModSlugs []string
}

type ExportedProfile struct {
	name     string
	modSlugs []mod.Slug
}

func ReformExportedProfile(dto ExportedProfileDto) ExportedProfile {
	modSlugs := make([]mod.Slug, len(dto.ModSlugs))
	for i, slug := range dto.ModSlugs {
		modSlugs[i] = mod.ReformSlugFromString(slug)
	}

	return ExportedProfile{
		name:     dto.Name,
		modSlugs: modSlugs,
	}
}

func (ep ExportedProfile) Name() string {
	return ep.name
}

func (ep ExportedProfile) ModSlugs() []mod.Slug {
	return ep.modSlugs
}
