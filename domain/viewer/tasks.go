package viewer

type Task string

const (
	TaskQuit           Task = "quit"
	TaskFirstTimeSetup Task = "first_time_setup"
	TaskNewProfile     Task = "new_profile"
	TaskSwitchProfile  Task = "switch_profile"
	TaskDeleteProfile  Task = "delete_profile"
	TaskExportProfile  Task = "export_profile"
	TaskImportProfile  Task = "import_profile"
	TaskSearchTerm     Task = "search_term"
	TaskAddMod         Task = "add_mod"
	TaskRemoveMod      Task = "remove_mod"
	TaskUpdateMods     Task = "update_mods"
)
