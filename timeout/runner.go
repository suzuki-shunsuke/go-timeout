package timeout

import (
	"context"
	"os/exec"
	"syscall"
	"time"
)

type (
	Runner struct {
		killAfter   time.Duration
		firstSignal syscall.Signal
		sig         chan syscall.Signal
	}
)

func NewRunner(killAfter time.Duration) *Runner {
	return &Runner{
		killAfter:   killAfter,
		firstSignal: syscall.SIGINT,
		sig:         make(chan syscall.Signal, 1),
	}
}

func (runner *Runner) SendSignal(sig syscall.Signal) {
	runner.sig <- sig
}

func (runner *Runner) SetSignal(sig syscall.Signal) {
	runner.firstSignal = sig
}

func (runner *Runner) Run(ctx context.Context, cmd *exec.Cmd) error {
	if err := cmd.Start(); err != nil {
		return &StartError{err: err}
	}
	targetID := 0
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err != nil {
		targetID = cmd.Process.Pid
	} else {
		targetID = -pgid
	}

	exitChan := make(chan error, 1)
	killAfterChan := make(chan struct{}, 1)

	go func() {
		exitChan <- cmd.Wait()
	}()

	for {
		select {
		case <-killAfterChan:
			if err := syscall.Kill(targetID, syscall.SIGKILL); err != nil {
				return &KillError{err: err, id: targetID}
			}
		case sig := <-runner.sig:
			if err := syscall.Kill(targetID, sig); err != nil {
				return &KillError{err: err, id: targetID}
			}
			if runner.killAfter > 0 {
				time.AfterFunc(runner.killAfter, func() {
					killAfterChan <- struct{}{}
				})
			}
		case <-ctx.Done():
			if err := syscall.Kill(targetID, runner.firstSignal); err != nil {
				return &KillError{err: err, id: targetID}
			}
			if runner.killAfter > 0 {
				time.AfterFunc(runner.killAfter, func() {
					killAfterChan <- struct{}{}
				})
			}
		case err := <-exitChan:
			if err == nil {
				return nil
			}
			return &WaitError{err: err}
		}
	}
}
