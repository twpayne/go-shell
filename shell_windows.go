// +build windows

package shell

import (
	"errors"
	"os"
	"os/user"
)

func UserShell(u *user.User) (string, error) {
	comSpecEnv := os.Getenv("ComSpec")
	if comSpecEnv == "" {
		return "", errors.New("Could not find shell")
	}
	return comSpecEnv, nil
}
