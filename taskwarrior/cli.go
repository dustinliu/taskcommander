package warrior

import "os/exec"

func Taskwarrior(args ...string) ([]byte, error) {
	command := exec.Command(taskCmd, args...)
	return command.Output()
}
