package viewer

type Task string

const (
	TaskQuit            Task = "quit"
	TaskFirstTimeSetup  Task = "first_time_setup"
	TaskNewProfile      Task = "new_profile"
	TaskSwitchProfile   Task = "switch_profile"
	TaskDeleteProfile   Task = "delete_profile"
	TaskSearchTerm      Task = "search_term"
	TaskAddModToProfile Task = "add_mod_to_profile"
)
