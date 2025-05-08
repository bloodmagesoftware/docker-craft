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
	exe, err := os.Executable()
	if err != nil {
		exe = os.Args[0]
	}
	targetDir := filepath.Join(u.HomeDir, ".docker", "cli-plugins", filepath.Base(os.Args[0]))
	if _, err := os.Stat(targetDir); err == nil {
		_ = os.Remove(targetDir)
	}
	if err := os.Symlink(exe, targetDir); err != nil {
		fmt.Fprintf(os.Stderr, "Error linking file: %s\n", err)
		os.Exit(1)
	}
}
