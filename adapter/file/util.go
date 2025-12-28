package file

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Model interface {
	profileModel | modModel
}

func read[M Model](fileName string) (map[string]M, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			newMap := make(map[string]M)
			writeErr := write(fileName, newMap)
			if writeErr != nil {
				return nil, writeErr
			}

			return newMap, nil
		}

		return nil, err
	}

	models := make(map[string]M)
	err = json.Unmarshal(bytes, &models)
	if err != nil {
		return nil, err
	}

	return models, nil
}

func readAllInDir[M Model](dirName string) (map[string]M, error) {
	entries, err := os.ReadDir(dirName)
	if err != nil {
		return nil, err
	}

	models := make(map[string]M)
	for _, entry := range entries {
		if !entry.Type().IsRegular() {
			continue
		}

		fileModels, err := read[M](entry.Name())
		if err != nil {
			return nil, err
		}

		models = mergeMaps(models, fileModels)
	}

	return models, nil
}

func write[M Model](fileName string, models map[string]M) error {
	bytes, err := json.MarshalIndent(models, "", "    ")
	if err != nil {
		return err
	}

	dir, name := filepath.Split(fileName)
	if strings.TrimSpace(dir) != "" {
		dir, err = filepath.Abs(dir)
		if err != nil {
			return err
		}

		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(filepath.Join(dir, name))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func mergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	if len(maps) == 0 {
		return nil
	}

	if len(maps) == 1 {
		return maps[0]
	}

	combined := make(map[K]V)
	for _, treasureMap := range maps {
		for key, value := range treasureMap {
			combined[key] = value
		}
	}

	return combined
}
