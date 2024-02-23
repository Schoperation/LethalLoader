package option

import "strings"

type CmdName string

const (
	TaskQuit           CmdName = "quit_task"
	TaskFirstTimeSetup CmdName = "first_time_setup_task"
	TaskCreateProfile  CmdName = "create_profile_task"
	TaskSwitchProfile  CmdName = "switch_profile_task"
	TaskDeleteProfile  CmdName = "delete_profile_task"
)

const (
	PageMainMenu      CmdName = "main_menu_page"
	PageProfileViewer CmdName = "profile_viewer_page"
)

func (cmdName CmdName) IsPage() bool {
	return strings.HasSuffix(string(cmdName), "_page")
}

func (cmdName CmdName) IsTask() bool {
	return strings.HasSuffix(string(cmdName), "_task")
}

func (cmdName CmdName) IsQuit() bool {
	return cmdName == TaskQuit
}
