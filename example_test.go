package shell_test

import (
	"fmt"
	"os/user"

	"github.com/twpayne/go-shell"
)

func Example_shell_CurrentUserShell() {
	shell, err := shell.CurrentUserShell()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(shell)
}

func Example_shell_UserShell() {
	u, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}
	s, err := shell.UserShell(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
