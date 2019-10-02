package shell_test

import (
	"fmt"
	"os/user"

	"github.com/twpayne/go-shell"
)

func Example_shell_CurrentUserShell() {
	shell, ok := shell.CurrentUserShell()
	_ = ok
	fmt.Println(shell)
}

func Example_shell_UserShell() {
	u, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}
	shell, ok := shell.UserShell(u)
	_ = ok
	fmt.Println(shell)
}
