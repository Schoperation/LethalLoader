package file

import (
	"encoding/json"
	"errors"
	"os"
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

func write[M Model](fileName string, models map[string]M) error {
	bytes, err := json.MarshalIndent(models, "", "    ")
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
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
