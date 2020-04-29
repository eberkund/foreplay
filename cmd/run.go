package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"foreplay/config"

	"github.com/fatih/color"
	"github.com/k0kubun/go-ansi"
	"github.com/mattn/go-isatty"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [hook]",
	Short: "Run hooks.",
	Args:  cobra.NoArgs,
	Run:   runRun,
}

type hookJob struct {
	config.Hook
	success *bool
	ticks   int
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func createHookJobs() []*hookJob {
	c, err := config.Get()
	if err != nil {
		panic(err)
	}
	var jobs []*hookJob
	for _, v := range c.Hooks {
		jobs = append(jobs, &hookJob{Hook: v})
	}
	return jobs
}

func runRun(cmd *cobra.Command, args []string) {
	jobs := createHookJobs()

	ansi.CursorHide()
	defer ansi.CursorShow()

	err := run(jobs)
	refresh(jobs, false)

	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func run(hooks []*hookJob) error {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	tick := time.NewTicker(125 * time.Millisecond)
	done := make(chan error)
	group := errgroup.Group{}

	for _, hook := range hooks {
		hook := hook
		group.Go(func() error {
			_, err := runHook(hook)
			success := err == nil
			hook.success = &success
			return err
		})
	}
	go func() {
		done <- group.Wait()
	}()
	for {
		select {
		case <-tick.C:
			refresh(hooks, true)
		case <-exit:
			return errors.New("user exited")
		case err := <-done:
			return err
		}
	}
}

func runHook(hook *hookJob) ([]byte, error) {
	cmd := exec.Command("sh")
	cmd.Stdin = bytes.NewBuffer([]byte(hook.Run))
	return cmd.CombinedOutput()
}

func (h hookJob) progressChar() string {
	charSet := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	successSymbol := "✓"
	errorSymbol := "✗"
	if h.success == nil {
		return charSet[h.ticks%len(charSet)]
	}
	if *h.success {
		return color.GreenString(successSymbol)
	}
	return color.RedString(errorSymbol)
}

func refresh(hooks []*hookJob, reset bool) {
	var writer io.Writer
	if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		writer = os.Stdout
	} else {
		writer = ansi.NewAnsiStdout()
	}
	table := tablewriter.NewWriter(writer)
	for _, v := range hooks {
		v.ticks++
		table.Append([]string{
			v.ID,
			v.progressChar(),
		})
	}
	table.Render()
	if reset {
		ansi.CursorPreviousLine(len(hooks) + 2)
	}
}
