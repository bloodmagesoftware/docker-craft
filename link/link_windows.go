//go:build windows
// +build windows

package link

import (
	"fmt"
	"os"
	"path/filepath"
)

func Link() {
	var targetDir string
	if up, ok := os.LookupEnv("USERPROFILE"); ok {
		targetDir = filepath.Join(up, ".docker", "cli-plugins", filepath.Base(os.Args[0]))
	} else {
		targetDir = filepath.Join("C:", "ProgramData", "Docker", "cli-plugins", filepath.Base(os.Args[0]))
	}
	exe, err := os.Executable()
	if err != nil {
		exe = os.Args[0]
	}
	if _, err := os.Stat(targetDir); err == nil {
		_ = os.Remove(targetDir)
	}
	if err := os.Symlink(exe, targetDir); err != nil {
		fmt.Fprintf(os.Stderr, "Error linking file: %s\n", err)
		os.Exit(1)
	}
}
