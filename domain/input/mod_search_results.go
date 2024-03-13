package input

import "schoperation/lethalloader/domain/profile"

type ModSearchResultsPageInput struct {
	Profile profile.Profile
	Term    string
}
