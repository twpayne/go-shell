package main

import (
	"fmt"
	"os"

	"github.com/twpayne/go-shell"
)

func main() {
	currentUserShell, ok := shell.CurrentUserShell()
	fmt.Println(currentUserShell)

	if !ok {
		os.Exit(1)
	}
}
