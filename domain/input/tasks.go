package input

import (
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
)

type SearchTermTaskInput struct {
	Profile         profile.Profile
	SkipCacheSearch bool
}

type AddModTaskInput struct {
	CachedMod    mod.Mod
	SearchResult mod.SearchResult
	Profile      profile.Profile
	UseCachedMod bool
}

type RemoveModTaskInput struct {
	Mod     mod.Mod
	Profile profile.Profile
}

type UpdateModsTaskInput struct {
	Listings []mod.Listing
	Profile  profile.Profile
}

type SwitchProfileTaskInput struct {
	OldProfile profile.Profile
	NewProfile profile.Profile
}
