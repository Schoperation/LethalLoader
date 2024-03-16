package input

import "schoperation/lethalloader/domain/profile"

type SearchTermTaskInput struct {
	Profile         profile.Profile
	SkipCacheSearch bool
}
