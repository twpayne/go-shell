package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/twpayne/go-shell"
)

func run() error {
	u, err := user.Current()
	if err != nil {
		return err
	}
	shell, err := shell.UserShell(u)
	if err != nil {
		return err
	}
	fmt.Println(shell)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
