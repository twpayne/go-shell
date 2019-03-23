// +build windows

package shell

import (
	"errors"
	"os"
)

// CurrentUserShell returns the current user's shell.
func CurrentUserShell() (string, error) {
	comSpecEnv := os.Getenv("ComSpec")
	if comSpecEnv == "" {
		return "", errors.New("could not find shell")
	}
	return comSpecEnv, nil
}
