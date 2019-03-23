// +build !darwin,!windows

package shell

import (
	"os/exec"
	"os/user"
	"strings"
)

// UserShell returns u's shell.
func UserShell(u *user.User) (string, error) {
	output, err := exec.Command("getent", "passwd", u.Username).Output()
	if err != nil {
		return "", err
	}
	fields := strings.Split(strings.TrimSuffix(string(output), "\n"), ":")
	if len(fields) < 7 {
		return "", parseError(output)
	}
	return fields[6], nil
}
