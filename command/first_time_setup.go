package command

type FirstTimeSetupCommand struct {
}

func NewFirstTimeSetupCommand() FirstTimeSetupCommand {
	return FirstTimeSetupCommand{}
}

func (cmd FirstTimeSetupCommand) Run() error {
	return nil
}
