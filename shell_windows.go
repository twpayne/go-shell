// +build windows

package shell

import (
	"errors"
	"os"
	"os/user"
)

func UserShell(u *user.User) (string, error) {
	ComSpecEnv := os.Getenv("ComSpec")
	if ComSpecEnv == "" {
		return "", errors.New("Could not find shell")
	}
	return ComSpecEnv, nil
}
