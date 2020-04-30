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
	Shell   *exec.Cmd
	Printer common.Registerable
	Hooks   []config.Hook
}

func (r Run) Start() {
	if skip() {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	var hookErr error
	results := make(chan common.Result)
	cleanup := r.Printer.Register(ctx, r.Hooks, results)
	group := errgroup.Group{}

	for _, hook := range r.Hooks {
		hook := hook
		group.Go(func() error {
			r.Shell.Stdin = bytes.NewBuffer([]byte(hook.Run))
			out, err := r.Shell.Output()
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

func skip() bool {
	foreplaySkipHooks := os.Getenv("FOREPLAY_SKIP_HOOKS")
	skip, _ := strconv.ParseBool(foreplaySkipHooks)
	return skip
}
