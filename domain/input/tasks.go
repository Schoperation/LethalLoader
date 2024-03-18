package input

import (
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
)

type SearchTermTaskInput struct {
	Profile         profile.Profile
	SkipCacheSearch bool
}

type AddModToProfileTaskInput struct {
	CachedMod    mod.Mod
	SearchResult mod.SearchResult
	Profile      profile.Profile
	UseCachedMod bool
}

type RemoveModFromProfileTaskInput struct {
	Mod     mod.Mod
	Profile profile.Profile
}
