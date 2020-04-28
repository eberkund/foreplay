package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"

	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
	"gopkg.in/yaml.v2"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [hook]",
	Short: "Run hooks.",
	Args:  cobra.MaximumNArgs(1),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runRun,
}

type Config struct {
	Hooks []Hook
}

type Hook struct {
	Id         string
	Command    string
	Args       []string
	WorkingDir string
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, args []string) {
	fmt.Println("run called")
	var c Config
	data, err := ioutil.ReadFile(".foreplay.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}

	wd, _ := os.Getwd()
	fmt.Println(wd)

	var wg sync.WaitGroup
	progress := mpb.New(
		mpb.WithWaitGroup(&wg),
		mpb.WithWidth(40),
	)
	wg.Add(len(c.Hooks))

	for _, v := range c.Hooks {
		go func(v Hook) {
			spinner := progress.AddSpinner(
				int64(len(c.Hooks)),
				mpb.SpinnerOnLeft,
				mpb.SpinnerStyle([]string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"}),
				mpb.PrependDecorators(
					decor.Name(v.Id),
				),
				mpb.AppendDecorators(
					decor.OnComplete(
						decor.Elapsed(decor.ET_STYLE_GO), "done",
					),
				),
			)

			cmd := exec.Command(v.Command, v.Args...)
			err := cmd.Run()
			//out, err := cmd.CombinedOutput()
			//fmt.Println(string(out))

			if err != nil {
				//fmt.Println("problem encountered: ", err)
				os.Exit(1)
			}

			spinner.SetTotal(1, true)
			wg.Done()
		}(v)
	}

	progress.Wait()
}
