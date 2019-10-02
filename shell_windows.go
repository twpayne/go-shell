// +build windows

package shell

import (
	"os"
)

const defaultShell = "PowerShell.exe"

// CurrentUserShell returns the current user's shell.
func CurrentUserShell() (string, bool) {
	// If the ComSpec environment variable is set, use it.
	if comSpec := os.Getenv("ComSpec"); comSpec != "" {
		return comSpec, true
	}

	// Fallback to the default shell.
	return defaultShell, false
}
