package profile

import (
	"fmt"
	"schoperation/lethalloader/domain/mod"
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

	return mods
}

func (pf *Profile) Dto() ProfileDto {
	modDtos := make([]mod.ModDto, len(pf.mods))
	slugs := make([]string, len(pf.mods))
	i := 0
	for _, mod := range pf.mods {
		modDtos[i] = mod.Dto()
		slugs[i] = mod.Author() + "-" + mod.Name() + "-" + mod.Version()
	}

	return ProfileDto{
		Name:     pf.name,
		ModSlugs: slugs,
		Mods:     modDtos,
	}
}

func (pf *Profile) AddMod(newMod mod.Mod) error {
	if _, exists := pf.mods[newMod.Name()]; exists {
		return fmt.Errorf("mod is already added")
	}

	pf.mods[newMod.Name()] = newMod
	return nil
}

func (pf *Profile) RemoveMod(modToRemove mod.Mod) error {
	delete(pf.mods, modToRemove.Name())
	return nil
}
