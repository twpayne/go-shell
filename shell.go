// Package shell returns a user's shell across multiple platforms.
package shell

import (
	"fmt"
	"os/user"
)

type parseError []byte

func (pe parseError) Error() string {
	return fmt.Sprintf("shell: cannot parse %q", pe)
}

// CurrentUserShell is a convenience function that returns the current user's
// shell.
func CurrentUserShell() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return UserShell(currentUser)
}
