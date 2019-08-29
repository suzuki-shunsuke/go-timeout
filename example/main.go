package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/suzuki-shunsuke/go-timeout/timeout"
)

func main() {
	cmd := exec.Command("sh", "-c", `
trap "echo hello sigint; exit 1" SIGINT
trap "echo hello err; exit 1" ERR
trap "echo hello exit; exit 1" EXIT
trap "echo hello sigterm; exit 1" SIGTERM
echo sleep
ls | fzf
sleep 3600
`)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan, syscall.SIGHUP, syscall.SIGINT,
		syscall.SIGTERM, syscall.SIGQUIT)

	runner := timeout.NewRunner(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sentSignals := map[os.Signal]struct{}{}
	exitChan := make(chan error, 1)

	go func() {
		exitChan <- runner.Run(ctx, cmd)
	}()

	for {
		select {
		case err := <-exitChan:
			fmt.Println("command exit", err)
			return
		case sig := <-signalChan:
			if _, ok := sentSignals[sig]; ok {
				continue
			}
			sentSignals[sig] = struct{}{}
			runner.SendSignal(sig.(syscall.Signal))
		}
	}
}
