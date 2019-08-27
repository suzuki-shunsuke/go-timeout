package timeout

import (
	"context"
	"os/exec"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRunner(t *testing.T) {
	assert.NotNil(t, NewRunner(0))
}

func TestRunner_SendSignal(t *testing.T) {
	runner := NewRunner(0)
	runner.SendSignal(syscall.SIGINT)
	a := <-runner.sig
	assert.Equal(t, syscall.SIGINT, a)
}

func TestRunner_SetSignal(t *testing.T) {
	runner := NewRunner(0)
	runner.SetSignal(syscall.SIGINT)
	assert.Equal(t, syscall.SIGINT, runner.firstSignal)
}

func TestRunner_Run(t *testing.T) {
	runner := NewRunner(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	assert.Nil(t, runner.Run(ctx, exec.Command("true")))
	assert.NotNil(t, runner.Run(ctx, exec.Command("false")))
}
