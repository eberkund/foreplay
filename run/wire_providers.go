package run

import (
	"context"
	"os"
	"os/exec"

	"foreplay/config"
	"foreplay/output"
	"foreplay/output/common"
)

func GetConfig() (config.Config, error) {
	return config.Get()
}

func GetShell() *exec.Cmd {
	return exec.CommandContext(context.Background(), "sh")
}

func GetPrinter(c config.Config) common.Registerable {
	return output.GetOutput(c.Style)
}

func GetExit() func(int) {
	return os.Exit
}

func GetRun(
	cmd *exec.Cmd,
	c config.Config,
	printer common.Registerable,
	exit func(int),
) *Run {
	return &Run{
		Shell:   cmd,
		Printer: printer,
		Hooks:   c.Hooks,
		exit:    exit,
	}
}
