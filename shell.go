// Package shell returns a user's shell across multiple platforms.

// +build !windows

package shell

import (
	"os/user"
)

// CurrentUserShell is a convenience function that returns the current user's
// shell.
func CurrentUserShell() (string, bool) {
	currentUser, err := user.Current()
	if err != nil {
		return defaultShell, false
	}
	return UserShell(currentUser)
}
