// +build !darwin,!windows

package shell

import (
	"bufio"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

const defaultShell = "/bin/sh"

// CurrentUserShell returns the current user's shell.
func CurrentUserShell() (string, bool) {
	if u, err := user.Current(); err == nil {
		// If getpwnam_r is available, use it.
		if shell, ok := cgoGetUserShell(u.Username); ok {
			return shell, true
		}

		// If getent is available, use it.
		if getent, err := exec.LookPath("getent"); err == nil {
			if output, err := exec.Command(getent, "passwd", u.Username).Output(); err == nil {
				if fields := strings.SplitN(strings.TrimSuffix(string(output), "\n"), ":", 7); len(fields) == 7 {
					return fields[6], true
				}
			}
		}

		// If the user has an entry in /etc/passwd, use it.
		if f, err := os.Open("/etc/passwd"); err == nil {
			defer f.Close()
			s := bufio.NewScanner(f)
			for s.Scan() {
				if fields := strings.SplitN(strings.TrimSuffix(s.Text(), "\n"), ":", 7); len(fields) == 7 && fields[0] == u.Username {
					return fields[6], true
				}
			}
			_ = s.Err() // Ignore errors.
		}
	}

	// If the SHELL environment variable is set, use it.
	if shell, ok := os.LookupEnv("SHELL"); ok {
		return shell, true
	}

	// Fallback to the default shell.
	return defaultShell, false
}
