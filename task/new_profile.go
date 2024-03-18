package task

import (
	"bufio"
	"fmt"
	"os"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
	"strings"
)

type newProfileSaver interface {
	GetAll() ([]profile.Profile, error)
	Save(pf profile.Profile) error
}

type NewProfileTask struct {
	newProfileSaver newProfileSaver
}

func NewNewProfileTask(
	newProfileSaver newProfileSaver,
) NewProfileTask {
	return NewProfileTask{
		newProfileSaver: newProfileSaver,
	}
}

func (task NewProfileTask) Do(args any) (viewer.TaskResult, error) {
	existingProfiles, err := task.newProfileSaver.GetAll()
	if err != nil {
		return viewer.TaskResult{}, err
	}

	existingProfileNames := make(map[string]bool, len(existingProfiles))
	for _, pf := range existingProfiles {
		existingProfileNames[strings.ToLower(pf.Name())] = true
	}

	fmt.Printf("\n")
	fmt.Printf("Name?\n")

	newProfileName := ""
	reader := bufio.NewReader(os.Stdin)
	for {
		newProfileName, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("The hell was that?\n")
			continue
		}

		newProfileName = strings.TrimSuffix(newProfileName, "\n")

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
		return viewer.TaskResult{}, err
	}

	err = task.newProfileSaver.Save(newProfile)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	return viewer.NewTaskResult(viewer.PageProfileViewer, newProfile), nil
}
