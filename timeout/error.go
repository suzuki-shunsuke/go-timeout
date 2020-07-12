package timeout

import (
	"syscall"
)

type (
	StartError struct {
		err error
	}

	WaitError struct {
		err error
	}

	KillError struct {
		err error
		id  int
		sig syscall.Signal
	}
)

func (se StartError) Error() string {
	return se.err.Error()
}

func (se StartError) Cause() error {
	return se.err
}

func (we WaitError) Error() string {
	return we.err.Error()
}

func (we WaitError) Cause() error {
	return we.err
}

func (ke KillError) Error() string {
	return ke.err.Error()
}

func (ke KillError) Cause() error {
	return ke.err
}

func (ke KillError) Signal() syscall.Signal {
	return ke.sig
}

func (ke KillError) PID() int {
	return ke.id
}
