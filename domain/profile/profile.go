package profile

import (
	"fmt"
	"schoperation/lethalloader/domain/mod"
	"strings"
)

type ProfileDto struct {
	Name string
	Mods []mod.ModDto
}

type Profile struct {
	name string
	mods []mod.Mod
}

func NewBlankProfile(dto ProfileDto) (Profile, error) {
	if strings.TrimSpace(dto.Name) == "" {
		return Profile{}, fmt.Errorf("profile must have name")
	}

	return Profile{
		name: dto.Name,
		mods: []mod.Mod{},
	}, nil
}

func ReformProfile(dto ProfileDto) Profile {
	mods := make([]mod.Mod, len(dto.Mods))
	for i, modDto := range dto.Mods {
		mods[i] = mod.ReformMod(modDto)
	}

	return Profile{
		name: dto.Name,
		mods: mods,
	}
}

func (pf Profile) Name() string {
	return pf.name
}

func (pf Profile) Mod() []mod.Mod {
	return pf.mods
}
