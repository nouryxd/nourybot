package common

var (
	tempCommands = 0
)

// CommandUsed is called on every command incremenenting tempCommands.
func CommandUsed() {
	tempCommands++
}

// GetCommandsUsed returns the amount of commands that have been used
// since the last restart.
func GetCommandsUsed() int {
	return tempCommands
}
