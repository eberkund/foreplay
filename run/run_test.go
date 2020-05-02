package run

import (
	"bytes"
	"testing"

	"foreplay/config"
	"foreplay/output/plain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunStart(t *testing.T) {
	var buf bytes.Buffer
	shell, err := GetShell()
	require.NoError(t, err)
	runner := GetRun(
		shell,
		config.Config{
			Hooks: []config.Hook{{
				ID:  "foo",
				Run: "pwd",
			}},
		},
		plain.New(&buf),
	)
	err = runner.Start()
	require.NoError(t, err)
	require.NotEmpty(t, buf.String())
}

func TestHookError(t *testing.T) {
	var buf bytes.Buffer
	shell, err := GetShell()
	require.NoError(t, err)
	runner := GetRun(
		shell,
		config.Config{
			Hooks: []config.Hook{{
				ID:  "bar",
				Run: "exit 1",
			}},
		},
		plain.New(&buf),
	)
	err = runner.Start()
	assert.Error(t, err)
}
