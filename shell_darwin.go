// +build darwin

package shell

import (
	"os"
	"os/exec"
	"os/user"
	"regexp"
)

const defaultShell = "/bin/bash"

//nolint:gochecknoglobals
var dsclUserShellRegexp = regexp.MustCompile(`\AUserShell:\s+(.*?)\s*\z`)

// CurrentUserShell returns the current user's shell.
func CurrentUserShell() (string, bool) {
	if u, err := user.Current(); err == nil {
		// If getpwnam_r is available, use it.
		if shell, ok := cgoGetUserShell(u.Username); ok {
			return shell, true
		}

		// If dscl is available, use it.
		//nolint:gosec
		if output, err := exec.Command("dscl", ".", "-read", u.HomeDir, "UserShell").Output(); err == nil {
			if m := dsclUserShellRegexp.FindSubmatch(output); m != nil {
				return string(m[1]), true
			}
		}
	}

	// If the SHELL environment variable is set, use it.
	if shell, ok := os.LookupEnv("SHELL"); ok {
		return shell, true
	}

	// Fallback to the default shell.
	return defaultShell, false
}
