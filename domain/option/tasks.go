package option

type TaskName string

const (
	TaskQuit           TaskName = "quit"
	TaskFirstTimeSetup TaskName = "first_time_setup"
	TaskNewProfile     TaskName = "new_profile"
	TaskSwitchProfile  TaskName = "switch_profile"
	TaskDeleteProfile  TaskName = "delete_profile"
)
