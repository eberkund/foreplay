package spinner

import (
	"bytes"
	"context"
	"os/exec"
	"testing"
	"time"

	"foreplay/config"
	"foreplay/output/common"

	"github.com/stretchr/testify/require"
)

func TestSpinnersOutput(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	results := make(chan common.Result)
	var buf bytes.Buffer
	p := New(&buf)
	done := p.Register(ctx, []config.Hook{}, results)
	cancel()
	<-done
	require.Empty(t, buf.String())
}

func TestSpinnerPrintsResults(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	results := make(chan common.Result)
	var buf bytes.Buffer
	p := New(&buf)
	hooks := []config.Hook{
		{ID: "foo"},
		{ID: "bar"},
	}
	done := p.Register(ctx, hooks, results)
	go func() {
		time.Sleep(125 * time.Millisecond)
		results <- common.Result{
			Hook: config.Hook{
				ID: "foo",
			},
			Err: nil,
			Out: []byte("hello world"),
		}
		results <- common.Result{
			Hook: config.Hook{
				ID: "bar",
			},
			Err: &exec.ExitError{},
			Out: nil,
		}
		cancel()
	}()
	<-done

	require.Contains(t, buf.String(), `| foo | ⣽ |`)
	require.Contains(t, buf.String(), `| bar | ⣽ |`)

	require.Contains(t, buf.String(), `| foo | ✓ |`)
	require.Contains(t, buf.String(), `| bar | ✗ |`)
}
