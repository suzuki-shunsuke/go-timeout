package timeout

import (
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunner_SendSignal(t *testing.T) { //nolint:paralleltest
	runner := NewRunner(0)
	runner.SendSignal(syscall.SIGINT)
	a := <-runner.sig
	assert.Equal(t, syscall.SIGINT, a)
}

func TestRunner_SetSignal(t *testing.T) { //nolint:paralleltest
	runner := NewRunner(0)
	runner.SetSignal(syscall.SIGINT)
	assert.Equal(t, syscall.SIGINT, runner.firstSignal)
}
