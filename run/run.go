package run

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"

	"foreplay/config"
	"foreplay/output/common"

	"golang.org/x/sync/errgroup"
)

type Run struct {
	Shell   string
	Printer common.Registerable
	Hooks   []config.Hook
}

// Start takes the list of Hooks and executes them.
func (r Run) Start() (err error) {
	if skip() {
		return nil
	}
	ctx, cancel := context.WithCancel(context.Background())
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	results := make(chan common.Result)
	cleanup := r.Printer.Register(ctx, r.Hooks, results)
	group := errgroup.Group{}

	for _, hook := range r.Hooks {
		hook := hook
		group.Go(func() error {
			cmd := r.createCmd(hook.Run)
			out, err := cmd.Output()
			results <- common.Result{
				Hook: hook,
				Err:  err,
				Out:  out,
			}
			return err
		})
	}
	go func() {
		err = group.Wait()
		cancel()
	}()
	func() {
		for {
			select {
			case <-exit:
				err = errors.New("user exit")
				cancel()
			case <-ctx.Done():
				return
			}
		}
	}()
	<-cleanup
	return
}

func (r Run) createCmd(script string) *exec.Cmd {
	cmd := exec.Command(r.Shell)
	cmd.Stdin = bytes.NewBuffer([]byte(script))
	return cmd
}

func skip() bool {
	foreplaySkipHooks := os.Getenv("FOREPLAY_SKIP_HOOKS")
	skip, _ := strconv.ParseBool(foreplaySkipHooks)
	return skip
}
