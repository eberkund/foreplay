package run

import (
	"context"
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

func GetRun(
	cmd *exec.Cmd,
	c config.Config,
	printer common.Registerable,
) *Run {
	return &Run{
		Shell:   cmd,
		Printer: printer,
		Hooks:   c.Hooks,
	}
}
