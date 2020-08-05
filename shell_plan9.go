// +build plan9

package shell

const defaultShell = "/bin/rc"

// CurrentUserShell returns the current user's shell.
func CurrentUserShell() (string, bool) {
	return defaultShell, false
}
