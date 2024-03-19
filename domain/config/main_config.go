package config

type MainConfigDto struct {
	GameFilesPath   string
	SelectedProfile string
}

type MainConfig struct {
	gameFilesPath   string
	selectedProfile string
}

func ReformMainConfig(dto MainConfigDto) MainConfig {
	return MainConfig{
		gameFilesPath:   dto.GameFilesPath,
		selectedProfile: dto.SelectedProfile,
	}
}

func (config *MainConfig) GameFilesPath() string {
	return config.gameFilesPath
}

func (config *MainConfig) SelectedProfile() string {
	return config.selectedProfile
}

func (config *MainConfig) Dto() MainConfigDto {
	return MainConfigDto{
		GameFilesPath:   config.gameFilesPath,
		SelectedProfile: config.selectedProfile,
	}
}

func (config *MainConfig) UpdateGameFilesPath(newPath string) {
	config.gameFilesPath = newPath
}

func (config *MainConfig) UpdateSelectedProfile(profileName string) {
	config.selectedProfile = profileName
}
