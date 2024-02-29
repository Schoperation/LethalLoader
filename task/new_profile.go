package task

import (
	"fmt"
	"schoperation/lethalloader/domain/profile"
	"strings"
)

type allProfilesGetter interface {
	GetAll() ([]profile.Profile, error)
}

type NewProfileTask struct {
	allProfilesGetter allProfilesGetter
}

func NewNewProfileTask(
	allProfilesGetter allProfilesGetter,
) NewProfileTask {
	return NewProfileTask{
		allProfilesGetter: allProfilesGetter,
	}
}

func (task NewProfileTask) Do(args ...any) error {
	existingProfiles, err := task.allProfilesGetter.GetAll()
	if err != nil {
		return err
	}

	existingProfileNames := make(map[string]bool, len(existingProfiles))
	for _, pf := range existingProfiles {
		existingProfileNames[pf.Name()] = true
	}

	fmt.Printf("\n")
	fmt.Printf("Name?\n")

	newProfileName := ""
	for {
		fmt.Scanf("%s", &newProfileName)

		if strings.TrimSpace(newProfileName) != "" {
			if _, alreadyExists := existingProfileNames[newProfileName]; !alreadyExists {
				break
			}

			fmt.Printf("Profile name already exists.\n")
		}

		fmt.Printf("Bruh that ain't a real name...\n")
	}

	newProfile, err := profile.NewBlankProfile(profile.ProfileDto{
		Name: newProfileName,
	})
	if err != nil {
		return err
	}

	return nil
}
