package shell

import "testing"

func TestShell(t *testing.T) {
	if shell, ok := CurrentUserShell(); shell == "" || !ok {
		t.Errorf("CurrentUserShell() == %v, %v, want !\"\", true", shell, ok)
	}
}
