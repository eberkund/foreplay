package run

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"foreplay/config"
	"foreplay/output/plain"

	"github.com/stretchr/testify/require"
)

func TestRunStart(t *testing.T) {
	var buf bytes.Buffer
	cmd := exec.Command("sh")
	runner := GetRun(
		cmd,
		config.Config{
			Hooks: []config.Hook{{
				ID:  "foo",
				Run: "pwd",
			}},
		},
		plain.New(&buf),
		os.Exit,
	)
	runner.Start()
	require.NotEmpty(t, buf.String())
}

func TestHookError(t *testing.T) {
	var buf bytes.Buffer
	cmd := exec.Command("sh")

	var code int
	runner := GetRun(
		cmd,
		config.Config{
			Hooks: []config.Hook{{
				ID:  "bar",
				Run: "exit 1",
			}},
		},
		plain.New(&buf),
		func(c int) {
			code = c
		},
	)
	runner.Start()
	require.Equal(t, 1, code)
}
