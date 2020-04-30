package run

import (
	"bytes"
	"os/exec"
	"testing"

	"foreplay/config"
	"foreplay/output/plain"
)

func TestRunStart(t *testing.T) {
	var buf bytes.Buffer
	//done := make(chan interface{})
	//printer := &mockstest.Registerable{}
	//printer.On("Register", mock.Anything, mock.Anything, mock.Anything).Return(done)
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
}
