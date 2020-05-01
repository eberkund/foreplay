package run

import (
	"bytes"
	"os"
	"testing"

	"foreplay/config"
	"foreplay/output/plain"

	"github.com/stretchr/testify/require"
)

func TestRunStart(t *testing.T) {
	var buf bytes.Buffer
	runner := GetRun(
		"sh",
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

	var code int
	runner := GetRun(
		"sh",
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
