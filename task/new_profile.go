package task

import (
	"fmt"
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
	"strings"
)

type newProfileSaver interface {
	GetAll() ([]profile.Profile, error)
	Save(pf profile.Profile) error
}

type bepInExListingGetter interface {
	GetByNameAndAuthor(name, author string) (mod.Listing, error)
}

type bepInExGetter interface {
	GetByModListing(listing mod.Listing) (mod.Mod, error)
}

type NewProfileTask struct {
	newProfileSaver      newProfileSaver
	bepInExListingGetter bepInExListingGetter
	bepInExGetter        bepInExGetter
}

func NewNewProfileTask(
	newProfileSaver newProfileSaver,
	bepInExListingGetter bepInExListingGetter,
	bepInExGetter bepInExGetter,
) NewProfileTask {
	return NewProfileTask{
		newProfileSaver:      newProfileSaver,
		bepInExListingGetter: bepInExListingGetter,
		bepInExGetter:        bepInExGetter,
	}
}

func (task NewProfileTask) Do(args ...any) (any, error) {
	existingProfiles, err := task.newProfileSaver.GetAll()
	if err != nil {
		return nil, err
	}

	existingProfileNames := make(map[string]bool, len(existingProfiles))
	for _, pf := range existingProfiles {
		existingProfileNames[strings.ToLower(pf.Name())] = true
	}

	fmt.Printf("\n")
	fmt.Printf("Name?\n")

	newProfileName := ""
	for {
		fmt.Scanf("%s", &newProfileName)

		if strings.TrimSpace(newProfileName) != "" {
			if _, alreadyExists := existingProfileNames[strings.ToLower(newProfileName)]; !alreadyExists {
				break
			}

			fmt.Printf("Profile name already exists.\n")
			continue
		}

		fmt.Printf("Bruh that ain't a real name...\n")
	}

	newProfile, err := profile.NewBlankProfile(profile.ProfileDto{
		Name: newProfileName,
	})
	if err != nil {
		return nil, err
	}

	bepInExListing, err := task.bepInExListingGetter.GetByNameAndAuthor("BepInExPack", "BepInEx")
	if err != nil {
		return nil, err
	}

	bepInEx, err := task.bepInExGetter.GetByModListing(bepInExListing)
	if err != nil {
		return nil, err
	}

	err = newProfile.AddMod(bepInEx)
	if err != nil {
		return nil, err
	}

	err = task.newProfileSaver.Save(newProfile)
	if err != nil {
		return nil, err
	}

	return newProfile, nil
}
