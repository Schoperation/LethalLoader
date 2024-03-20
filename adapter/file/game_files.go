package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"schoperation/lethalloader/domain/mod"
	"schoperation/lethalloader/domain/profile"
)

type GameFilesDao struct {
}

func NewGameFilesDao() GameFilesDao {
	return GameFilesDao{}
}

func (dao GameFilesDao) slug(name, author, version string) string {
	return author + "-" + name + "-" + version
}

func (dao GameFilesDao) CheckDefaultPath() (string, error) {
	homeDir := os.Getenv("HOME")
	defaultGameFilePath := homeDir + "/.steam/steam/steamapps/common/Lethal Company"

	if runtime.GOOS == "windows" {
		defaultGameFilePath = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Lethal Company"
	}

	_, err := os.Stat(defaultGameFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}

		return "", err
	}

	return defaultGameFilePath, nil
}

func (dao GameFilesDao) CheckPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (dao GameFilesDao) AddFilesByMod(mod mod.ModDto, gameFilesPath string) error {
	modSlug := dao.slug(mod.Name, mod.Author, mod.Version)
	for _, file := range mod.Files {
		path := filepath.Join(gameFilesPath, file.Path)

		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}

		newFile, err := os.Create(filepath.Join(path, file.Name))
		if err != nil {
			return err
		}
		defer newFile.Close()

		cachedFile, err := os.Open(filepath.Join("modcache", modSlug, file.Path, file.Name))
		if err != nil {
			return err
		}
		defer cachedFile.Close()

		_, err = io.Copy(newFile, cachedFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dao GameFilesDao) DeleteFilesByMod(mod mod.ModDto, gameFilesPath string) error {
	for _, file := range mod.Files {
		path := filepath.Join(gameFilesPath, file.Path, file.Name)

		err := os.RemoveAll(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dao GameFilesDao) AddFilesByProfile(pf profile.ProfileDto, gameFilesPath string) error {
	modSlugs := make(map[string]bool)
	for _, slug := range pf.ModSlugs {
		modSlugs[slug] = true
	}

	for _, mod := range pf.Mods {
		modSlug := dao.slug(mod.Name, mod.Author, mod.Version)
		if _, exists := modSlugs[modSlug]; !exists {
			return fmt.Errorf("could not find mod within profile slugs")
		}

		err := dao.AddFilesByMod(mod, gameFilesPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dao GameFilesDao) DeleteFilesByProfile(pf profile.ProfileDto, gameFilesPath string) error {
	for _, mod := range pf.Mods {
		err := dao.DeleteFilesByMod(mod, gameFilesPath)
		if err != nil {
			return err
		}
	}

	return nil
}
