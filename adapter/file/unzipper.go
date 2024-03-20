package file

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"schoperation/lethalloader/domain/mod"
	"slices"
	"strings"
)

type FileUnzipper struct{}

func NewFileUnzipper() FileUnzipper {
	return FileUnzipper{}
}

func (fuz FileUnzipper) Unzip(zippedDto mod.FileDto) ([]mod.FileDto, error) {
	reader, err := zip.OpenReader(zippedDto.Path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	defer os.Remove(zippedDto.Path)

	err = os.MkdirAll("modcache"+fuz.sep()+zippedDto.Name, 0755)
	if err != nil {
		return nil, err
	}

	fileDtos := []mod.FileDto{}
	for _, f := range reader.File {
		fileDto, err := fuz.extractFile(f, "modcache"+fuz.sep()+zippedDto.Name)
		if err != nil {
			return nil, err
		}

		if fileDto.Name == "skip" {
			continue
		}

		fileDtos = append(fileDtos, fileDto)
	}

	return fileDtos, nil
}

func (fuz FileUnzipper) extractFile(file *zip.File, unzippedFolder string) (mod.FileDto, error) {
	rc, err := file.Open()
	if err != nil {
		return mod.FileDto{}, err
	}
	defer rc.Close()

	// author-modname-version
	names := strings.Split(unzippedFolder, "-")
	modName := names[1]

	// Remove extraneous modname folders.
	if file.Name == modName+fuz.sep() {
		return mod.FileDto{Name: "skip"}, nil
	}

	newPath := filepath.Join(unzippedFolder, file.Name)
	newPath = strings.Replace(newPath, modName+fuz.sep(), "", 1)
	newPath = fuz.fixPathForDlls(file, newPath, unzippedFolder)

	// Check for ZipSlip.
	if !strings.HasPrefix(newPath, filepath.Clean(unzippedFolder)+fuz.sep()) {
		return mod.FileDto{}, fmt.Errorf("illegal file path: %s", newPath)
	}

	if fuz.shouldSkipFile(file, newPath, unzippedFolder) {
		return mod.FileDto{Name: "skip"}, nil
	}

	err = os.MkdirAll(filepath.Dir(newPath), 0755)
	if err != nil {
		return mod.FileDto{}, err
	}

	f, err := os.Create(newPath)
	if err != nil {
		return mod.FileDto{}, err
	}
	defer f.Close()

	_, err = io.Copy(f, rc)
	if err != nil {
		return mod.FileDto{}, err
	}

	stats, err := f.Stat()
	if err != nil {
		return mod.FileDto{}, err
	}

	dtoPath := strings.TrimPrefix(newPath, unzippedFolder+fuz.sep())
	dtoPath = strings.TrimSuffix(dtoPath, stats.Name())

	return mod.FileDto{
		Name: stats.Name(),
		Path: dtoPath,
	}, nil
}

func (fuz FileUnzipper) shouldSkipFile(file *zip.File, path, unzippedFolder string) bool {
	if file.FileInfo().IsDir() {
		return true
	}

	// These file exceptions only apply to files in the base of the zip.
	if !fuz.isBasePath(path, file.FileInfo().Name(), unzippedFolder) {
		return false
	}

	skippedFiles := []string{
		"manifest.json",
		"icon.png",
	}

	if slices.Contains(skippedFiles, file.FileInfo().Name()) {
		return true
	}

	if strings.HasSuffix(file.FileInfo().Name(), ".md") {
		return true
	}

	return false
}

// Some mods have their DLLs and binaries in the wrong place:
//   - Base folder
//   - plugins folder without the parent BepInEx
//   - placeholder for another inevitable exception
//
// This function corrects the path for said files.
func (fuz FileUnzipper) fixPathForDlls(file *zip.File, path, unzippedFolder string) string {
	if !fuz.isBasePath(path, file.FileInfo().Name(), unzippedFolder) && strings.Contains(path, "BepInEx") {
		return path
	}

	if !strings.HasSuffix(file.FileInfo().Name(), ".dll") && strings.Contains(file.FileInfo().Name(), ".") {
		return path
	}

	// We all love BepInEx exceptions
	if file.FileInfo().Name() == "winhttp.dll" {
		return path
	}

	// Missing BepInEx; just add that to the path
	if !strings.Contains(path, "BepInEx"+fuz.sep()+"plugins") {
		return strings.Replace(path, "plugins", "BepInEx"+fuz.sep()+"plugins", 1)
	}

	return strings.Replace(path, file.Name, "BepInEx"+fuz.sep()+"plugins"+fuz.sep()+file.Name, 1)
}

func (fuz FileUnzipper) isBasePath(path, fileName, unzippedFolder string) bool {
	return strings.TrimSuffix(path, fuz.sep()+fileName) == unzippedFolder
}

// Goddamn I hate typing this out
func (fuz FileUnzipper) sep() string {
	return string(os.PathSeparator)
}
