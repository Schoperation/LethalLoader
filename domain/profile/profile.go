package profile

import (
	"fmt"
	"schoperation/lethalloader/domain/mod"
	"slices"
	"strings"
)

type ProfileDto struct {
	Name     string
	ModSlugs []string
	Mods     []mod.ModDto
}

type Profile struct {
	name string
	mods map[string]mod.Mod
}

func NewBlankProfile(dto ProfileDto) (Profile, error) {
	if strings.TrimSpace(dto.Name) == "" {
		return Profile{}, fmt.Errorf("profile must have name")
	}

	return Profile{
		name: dto.Name,
		mods: make(map[string]mod.Mod),
	}, nil
}

func ReformProfile(dto ProfileDto) Profile {
	mods := make(map[string]mod.Mod, len(dto.Mods))
	for _, modDto := range dto.Mods {
		mods[modDto.Name] = mod.ReformMod(modDto)
	}

	return Profile{
		name: dto.Name,
		mods: mods,
	}
}

func (pf *Profile) Name() string {
	return pf.name
}

func (pf *Profile) Mods() []mod.Mod {
	mods := make([]mod.Mod, len(pf.mods))
	i := 0
	for _, pfMod := range pf.mods {
		mods[i] = pfMod
		i++
	}

	slices.SortFunc(mods, func(a, b mod.Mod) int {
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

	return mods
}

func (pf *Profile) Mod(name string) (mod.Mod, error) {
	foundMod, exists := pf.mods[name]
	if !exists {
		return mod.Mod{}, fmt.Errorf("could not find mod %s", name)
	}

	return foundMod, nil
}

func (pf *Profile) Dto() ProfileDto {
	modDtos := make([]mod.ModDto, len(pf.mods))
	slugs := make([]string, len(pf.mods))
	i := 0
	for _, modd := range pf.mods {
		modDtos[i] = modd.Dto()
		slug := mod.ReformSlug(modd.Name(), modd.Author(), modd.Version())
		slugs[i] = slug.String()
		i++
	}

	return ProfileDto{
		Name:     pf.name,
		ModSlugs: slugs,
		Mods:     modDtos,
	}
}

func (pf *Profile) AddMod(newMod mod.Mod) {
	if _, exists := pf.mods[newMod.Name()]; exists {
		return
	}

	pf.mods[newMod.Name()] = newMod
}

func (pf *Profile) RemoveMod(modToRemove mod.Mod) {
	delete(pf.mods, modToRemove.Name())
}
