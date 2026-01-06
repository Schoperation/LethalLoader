package task

import (
	"fmt"
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
	"schoperation/lethalloader/domain/viewer"
)

type listingGetter interface {
	GetBySlug(slug mod.Slug) (mod.Listing, error)
}

type modDownloader interface {
	GetByModListing(listing mod.Listing) (mod.Mod, error)
}

type importedProfileSaver interface {
	Save(pf profile.Profile) error
}

type ImportProfileTask struct {
	listingGetter        listingGetter
	modDownloader        modDownloader
	importedProfileSaver importedProfileSaver
}

func NewImportProfileTask(
	listingGetter listingGetter,
	modDownloader modDownloader,
	importedProfileSaver importedProfileSaver,
) ImportProfileTask {
	return ImportProfileTask{
		listingGetter:        listingGetter,
		modDownloader:        modDownloader,
		importedProfileSaver: importedProfileSaver,
	}
}

func (task ImportProfileTask) Do(args any) (viewer.TaskResult, error) {
	pfToImport, ok := args.(profile.ExportedProfile)
	if !ok {
		return viewer.TaskResult{}, fmt.Errorf("could not cast profile")
	}

	fmt.Print("\n")
	fmt.Printf("Importing profile %s...\n", pfToImport.Name())

	newPf, err := profile.NewBlankProfile(profile.ProfileDto{
		Name: pfToImport.Name(),
	})
	if err != nil {
		return viewer.TaskResult{}, err
	}

	for _, slug := range pfToImport.ModSlugs() {
		fmt.Printf("\tDownloading %s v%s...\n", slug.Name(), slug.Version())

		listing, err := task.listingGetter.GetBySlug(slug)
		if err != nil {
			return viewer.TaskResult{}, err
		}

		mod, err := task.modDownloader.GetByModListing(listing)
		if err != nil {
			return viewer.TaskResult{}, err
		}

		newPf.AddMod(mod)
	}

	err = task.importedProfileSaver.Save(newPf)
	if err != nil {
		return viewer.TaskResult{}, err
	}

	return viewer.NewTaskResult(viewer.PageMainMenu, pfToImport), nil
}
