package service

import "os/exec"

func taskwarrior(args ...string) ([]byte, error) {
	command := exec.Command(taskCmd, args...)
	return command.Output()
}
