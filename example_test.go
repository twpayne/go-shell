package shell_test

import (
	"fmt"

	"github.com/twpayne/go-shell"
)

func Example_shell_CurrentUserShell() {
	shell, ok := shell.CurrentUserShell()
	_ = ok
	fmt.Println(shell)
}
