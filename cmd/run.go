package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
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
	Hooks []Hook `yaml:"hooks" jsonschema:"required"`
}

type Hook struct {
	ID  string `yaml:"id" jsonschema:"required"`
	Run string `yaml:"run" jsonschema:"required"`
}

type Result struct {
	Hook *Hook
	Err  error
	Out  string
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, args []string) {
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

	wg := sync.WaitGroup{}
	wg.Add(len(c.Hooks))

	s := spinner.New(spinner.CharSets[9], 125*time.Millisecond)
	s.HideCursor = true
	hookCh := make(chan *Result)

	for _, hook := range c.Hooks {
		hook := hook
		go func() {
			cmd := exec.Command("sh")
			cmd.Stdin = bytes.NewBuffer([]byte(hook.Run))
			out, err := cmd.CombinedOutput()
			hookCh <- &Result{
				Hook: &hook,
				Err:  err,
				Out:  string(out),
			}
			wg.Done()
		}()
	}
	s.Start()
	go func() {
		wg.Wait()
		time.Sleep(time.Second)
		close(hookCh)
	}()
	var results []*Result
	for r := range hookCh {
		results = append(results, r)
	}
	wg.Wait()
	s.Stop()

	table := tablewriter.NewWriter(os.Stdout)
	for _, r := range results {
		var status string
		if r.Err == nil {
			status = "✓"
		} else {
			status = "✗"
		}
		table.Append([]string{
			r.Hook.ID,
			status,
		})
	}
	table.Render()
	println()
}
