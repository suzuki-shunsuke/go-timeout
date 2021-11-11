//go:build !windows
// +build !windows

package timeout

import (
	"syscall"
)

func kill(pid int, sig syscall.Signal) error {
	return syscall.Kill(pid, sig) //nolint:wrapcheck
}

func getTargetID(pid int) int {
	pgid, err := syscall.Getpgid(pid)
	if err != nil {
		return pid
	}
	return -pgid
}
