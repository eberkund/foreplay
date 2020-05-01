package run

import (
	"bytes"
	"context"
	"errors"
	"io"
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
	exit    func(int)
	writer  io.Writer
}

func (r *Run) SetOut(out io.Writer) {
	r.writer = out
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
		r.exit(1)
	}
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
