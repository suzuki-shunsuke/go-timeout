package timeout_test

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/suzuki-shunsuke/go-timeout/timeout"
)

func Example() {
	cmd := exec.Command("true")
	runner := timeout.NewRunner(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := runner.Run(ctx, cmd); err != nil {
		fmt.Println(err)
	}
	// Output:
}
