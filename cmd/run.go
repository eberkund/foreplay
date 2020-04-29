package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/k0kubun/go-ansi"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v2"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [hook]",
	Short: "Run hooks.",
	Args:  cobra.NoArgs,
	Run:   runRun,
}

type Config struct {
	Hooks []*Hook `yaml:"hooks" jsonschema:"required"`
}

type Hook struct {
	ID      string `yaml:"id" jsonschema:"required"`
	Run     string `yaml:"run" jsonschema:"required"`
	success *bool
	count   int
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, args []string) {
	c := getConfig()

	ansi.CursorHide()
	defer ansi.CursorShow()

	err := run(c.Hooks)
	refresh(c.Hooks, false)

	if err != nil {
		os.Exit(1)
	}
}

func run(hooks []*Hook) error {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	tick := time.NewTicker(125 * time.Millisecond)

	done := make(chan interface{})
	group := errgroup.Group{}

	var hookErr error
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
		hookErr = group.Wait()
		close(done)
	}()
	for {
		select {
		case <-tick.C:
			refresh(hooks, true)
		case <-exit:
			return errors.New("user exited")
		case <-done:
			return hookErr
		}
	}
}

func getConfig() *Config {
	var c Config
	data, err := ioutil.ReadFile(".foreplay.yml")
	if err != nil {
		println("could not read config file")
		panic(err)
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		println("could not unmarshal config file")
		panic(err)
	}
	return &c
}

func runHook(hook *Hook) ([]byte, error) {
	cmd := exec.Command("sh")
	cmd.Stdin = bytes.NewBuffer([]byte(hook.Run))
	return cmd.CombinedOutput()
}

func (h Hook) progressChar() string {
	charSet := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	successSymbol := "✓"
	errorSymbol := "✗"
	if h.success == nil {
		return charSet[h.count%len(charSet)]
	}
	if *h.success {
		return successSymbol
	}
	return errorSymbol
}

func refresh(hooks []*Hook, reset bool) {
	for _, v := range hooks {
		v.count++
		fmt.Printf("%-020s %s\n", v.ID, v.progressChar())
	}
	if reset {
		ansi.CursorPreviousLine(len(hooks))
	}
}
