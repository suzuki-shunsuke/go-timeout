# go-timeout

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/go-timeout/timeout)
[![Build Status](https://cloud.drone.io/api/badges/suzuki-shunsuke/go-timeout/status.svg)](https://cloud.drone.io/suzuki-shunsuke/go-timeout)
[![Test Coverage](https://api.codeclimate.com/v1/badges/3d3b9a00cbd9a9188962/test_coverage)](https://codeclimate.com/github/suzuki-shunsuke/go-timeout/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-timeout)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-timeout)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-timeout.svg)](https://github.com/suzuki-shunsuke/go-timeout)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-timeout/main/LICENSE)

Golang library to treat the timeout of the external command.

This is inspired to [Songmu/timeout](https://github.com/Songmu/timeout).
Basically, we recommend to use [Songmu/timeout](https://github.com/Songmu/timeout) instead of this library.
We develop this library because we don't know how to run [fzf](https://github.com/junegunn/fzf) with [Songmu/timeout](https://github.com/Songmu/timeout).

## :warning: Deprecated. Go 1.20 has supported the feature officially

Please use [os/exec#CommandContext](https://pkg.go.dev/os/exec#CommandContext), `os/exec#Cmd.Cancel`, and `os/exec#Cmd.WaitDelay` instead.

https://pkg.go.dev/os/exec#Cmd

e.g.

```go
cmd := exec.CommandContext(ctx, exePath, args...)
setCancel(cmd)
if err := cmd.Run(); err != nil {
	return err
}

const waitDelay = 1000 * time.Hour

func setCancel(cmd *exec.Cmd) {
	cmd.Cancel = func() error {
		return cmd.Process.Signal(os.Interrupt) //nolint:wrapcheck
	}
	cmd.WaitDelay = waitDelay
}
```

## Example

[example/main.go](example/main.go)

## Reference

* https://github.com/suzuki-shunsuke/cmdx/issues/52

## Lience

[MIT](LICENSE)
