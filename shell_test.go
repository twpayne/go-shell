package shell_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-shell"
)

func TestCurrentUserShell(t *testing.T) {
	shell, ok := shell.CurrentUserShell()
	assert.True(t, ok)
	assert.NotEqual(t, "", shell)
}
