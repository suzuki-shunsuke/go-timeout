# go-timeout

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/go-timeout/timeout)
[![Build Status](https://cloud.drone.io/api/badges/suzuki-shunsuke/go-timeout/status.svg)](https://cloud.drone.io/suzuki-shunsuke/go-timeout)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/go-timeout/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/go-timeout)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-timeout)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-timeout)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-timeout.svg)](https://github.com/suzuki-shunsuke/go-timeout)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-timeout/master/LICENSE)

Golang library to treat the timeout of the external command.

This is inspired to [Songmu/timeout](https://github.com/Songmu/timeout).
Basically, we recommend to use [Songmu/timeout](https://github.com/Songmu/timeout) instead of this library.
We develop this library because we don't know how to run [fzf](https://github.com/junegunn/fzf) with [Songmu/timeout](https://github.com/Songmu/timeout).

## Example

[example/main.go](example/main.go)

## Reference

* https://github.com/suzuki-shunsuke/cmdx/issues/52

## Lience

[MIT](LICENSE)
