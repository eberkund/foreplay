package run

import (
	"bytes"
	"os"
	"testing"

	"github.com/eberkund/foreplay/config"
	"github.com/eberkund/foreplay/mockstest"
	"github.com/eberkund/foreplay/output/plain"

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
	require.Error(t, err)
}

func TestSkip(t *testing.T) {
	err := os.Setenv("FOREPLAY_SKIP", "true")
	require.NoError(t, err)

	shell, err := GetShell()
	require.NoError(t, err)

	m := mockstest.Registerable{}
	m.AssertNotCalled(t, "Run")

	runner := GetRun(
		shell,
		config.Config{
			Hooks: []config.Hook{{
				ID:  "bar",
				Run: "exit 1",
			}},
		},
		&m,
	)
	require.NoError(t, runner.Start())
}
