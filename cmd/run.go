package cmd

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"foreplay/config"
	"foreplay/output"
	"foreplay/output/common"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var runCmd = &cobra.Command{
	Use:   "run [hook]",
	Short: "Run hooks.",
	Args:  cobra.NoArgs,
	Run:   runRun,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	c, err := config.Get()
	if err != nil {
		panic(err)
	}

	var hookErr error
	results := make(chan common.Result)

	p := output.GetOutput(c.Style)
	cleanup := p.Register(ctx, c.Hooks, results)
	group := errgroup.Group{}

	for _, hook := range c.Hooks {
		hook := hook
		group.Go(func() error {
			out, err := runHook(ctx, hook)
			results <- common.Result{
				Hook: hook,
				Err:  err,
				Out:  out,
			}
			return err
		})
	}
	go func() {
		hookErr = group.Wait()
		cancel()
	}()
	func() {
		for {
			select {
			case <-exit:
				hookErr = errors.New("user exit")
				cancel()
			case <-ctx.Done():
				return
			}
		}
	}()
	<-cleanup
	if hookErr != nil {
		os.Exit(1)
	}
}

func runHook(ctx context.Context, hook config.Hook) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "sh")
	cmd.Stdin = bytes.NewBuffer([]byte(hook.Run))
	return cmd.Output()
}
