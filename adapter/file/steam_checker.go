package file

import (
	"os"
	"runtime"
)

type SteamChecker struct {
}

func NewSteamChecker() SteamChecker {
	return SteamChecker{}
}

func (checker SteamChecker) CheckDefault() (string, error) {
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

func (checker SteamChecker) Check(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
