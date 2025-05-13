package docker

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func ComposeUp(file string, d bool) error {
	args := []string{"compose", "-f", file, "up", "--build"}
	if d {
		args = append(args, "-d")
	}
	cmd := exec.Command("docker", args...)
	if !d {
		cmd.Stdin = os.Stdin
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if d {
		return cmd.Start()
	} else {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGTERM)

		defer func() {
			signal.Stop(sig)
			close(sig)
		}()

		go func() {
			defer recover()
			s := <-sig
			if cmd.Process != nil && cmd.Process.Pid > 0 {
				cmd.Process.Signal(s)
			}
		}()

		return cmd.Run()
	}
}
