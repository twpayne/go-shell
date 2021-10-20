package shell_test

import (
	"testing"

	"github.com/twpayne/go-shell"
)

func TestShell(t *testing.T) {
	if shell, ok := shell.CurrentUserShell(); shell == "" || !ok {
		t.Errorf("shell.CurrentUserShell() == %v, %v, want !\"\", true", shell, ok)
	}
}
