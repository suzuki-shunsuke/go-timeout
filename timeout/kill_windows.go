// +build windows

package timeout

import (
	"os/exec"
	"strconv"
	"syscall"
)

func kill(pid int, sig syscall.Signal) error {
	return exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(pid)).Run()
}

func getTargetID(pid int) int {
	return pid
}
