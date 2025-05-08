//go:build !windows
// +build !windows

package link

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func Link() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	if err := os.Symlink(
		os.Args[0],
		filepath.Join(u.HomeDir, ".docker", "cli-plugins", filepath.Base(os.Args[0])),
	); err != nil {
		fmt.Fprintf(os.Stderr, "Error linking file: %s\n", err)
		os.Exit(1)
	}
}
