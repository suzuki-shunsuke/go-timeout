package timeout_test

import (
	"context"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuki-shunsuke/go-timeout/timeout"
)

func TestNewRunner(t *testing.T) {
	assert.NotNil(t, timeout.NewRunner(0))
}

func TestRunner_Run(t *testing.T) {
	runner := timeout.NewRunner(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	assert.Nil(t, runner.Run(ctx, exec.Command("true")))
	assert.NotNil(t, runner.Run(ctx, exec.Command("false")))
}
