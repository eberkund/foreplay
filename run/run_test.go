package run

import (
	"bytes"
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
	)
	runner.Start()
	require.NotEmpty(t, buf.String())
}
