// +build windows

package shell

import (
	"errors"
	"os"
)

func CurrentUserShell() (string, error) {
	comSpecEnv := os.Getenv("ComSpec")
	if comSpecEnv == "" {
		return "", errors.New("Could not find shell")
	}
	return comSpecEnv, nil
}
