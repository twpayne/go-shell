package shell_test

import (
	"fmt"

	"github.com/twpayne/go-shell"
)

func Example_shell_CurrentUserShell() {
	shell, _ := shell.CurrentUserShell()
	fmt.Println(shell)
}
