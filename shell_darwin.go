// +build darwin

package shell

import (
	"os/exec"
	"os/user"
	"regexp"
)

var dsclUserShellRegexp = regexp.MustCompile(`\AUserShell:\s+(.*?)\s*\z`)

func UserShell(u *user.User) (string, error) {
	output, err := exec.Command("dscl", ".", "-read", u.HomeDir, "UserShell").Output()
	if err != nil {
		return "", err
	}
	m := dsclUserShellRegexp.FindSubmatch(output)
	if m == nil {
		return "", parseError(output)
	}
	return string(m[1]), nil
}
