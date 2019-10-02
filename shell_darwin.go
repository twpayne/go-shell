// +build darwin

package shell

import (
	"os"
	"os/exec"
	"os/user"
	"regexp"
)

const defaultShell = "/bin/bash"

var dsclUserShellRegexp = regexp.MustCompile(`\AUserShell:\s+(.*?)\s*\z`)

// UserShell returns u's shell.
func UserShell(u *user.User) (string, bool) {
	// If dscl is available, use it.
	if output, err := exec.Command("dscl", ".", "-read", u.HomeDir, "UserShell").Output(); err != nil {
		if m := dsclUserShellRegexp.FindSubmatch(output); m != nil {
			return string(m[1]), true
		}
	}

	// If the SHELL environment variable is set, use it.
	if shell, ok := os.LookupEnv("SHELL"); ok {
		return shell, true
	}

	// Fallback to the default shell.
	return defaultShell, false
}
