package run

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"foreplay/config"
	"foreplay/output/common"

	"golang.org/x/sync/errgroup"
)

type Run struct {
	Shell   string
	Printer common.Registerable
	Hooks   []config.Hook
	Timeout time.Duration
}

// Start takes the list of Hooks and executes them.
func (r Run) Start() error {
	if skip() {
		return nil
	}

	printerCtx, printerCancel := context.WithCancel(context.Background())
	execCtx, execCancel := r.createContext()

	results := make(chan common.Result)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		r.Printer.Run(printerCtx, r.Hooks, results)
		wg.Done()
	}()

	group := errgroup.Group{}
	for _, hook := range r.Hooks {
		hook := hook
		group.Go(func() error {
			cmd := r.createCmd(execCtx, hook.Run)
			out, err := cmd.Output()
			results <- common.Result{
				Hook: hook,
				Err:  err,
				Out:  out,
			}
			return err
		})
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-exit
		execCancel()
	}()

	err := group.Wait()
	printerCancel()
	wg.Wait()
	return err
}

func (r Run) createContext() (context.Context, context.CancelFunc) {
	if r.Timeout == 0 {
		return context.WithCancel(context.Background())
	}
	return context.WithTimeout(context.Background(), r.Timeout)
}

func (r Run) createCmd(ctx context.Context, script string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, r.Shell)
	cmd.Stdin = bytes.NewBuffer([]byte(script))
	return cmd
}

func skip() bool {
	foreplaySkipHooks := os.Getenv("FOREPLAY_SKIP")
	skip, _ := strconv.ParseBool(foreplaySkipHooks)
	return skip
}
