package config

type MainConfigDto struct {
	GameFilePath    string
	SelectedProfile string
}

type MainConfig struct {
	gameFilePath    string
	selectedProfile string
}

func ReformMainConfig(dto MainConfigDto) MainConfig {
	return MainConfig{
		gameFilePath:    dto.GameFilePath,
		selectedProfile: dto.SelectedProfile,
	}
}

func (config *MainConfig) GameFilePath() string {
	return config.gameFilePath
}

func (config *MainConfig) SelectedProfile() string {
	return config.selectedProfile
}

func (config *MainConfig) Dto() MainConfigDto {
	return MainConfigDto{
		GameFilePath:    config.gameFilePath,
		SelectedProfile: config.selectedProfile,
	}
}

func (config *MainConfig) UpdateGameFilePath(newPath string) error {
	return nil
}
