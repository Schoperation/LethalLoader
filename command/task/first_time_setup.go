package task

type FirstTimeSetupTask struct {
}

func NewFirstTimeSetupTask() FirstTimeSetupTask {
	return FirstTimeSetupTask{}
}

func (task FirstTimeSetupTask) Do() error {
	return nil
}
