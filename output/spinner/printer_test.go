package spinner

import (
	"bytes"
	"context"
	"os/exec"
	"sync"
	"testing"
	"time"

	"github.com/eberkund/foreplay/config"
	"github.com/eberkund/foreplay/output/common"

	"github.com/stretchr/testify/require"
)

func TestSpinnersOutput(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	results := make(chan common.Result)
	var buf bytes.Buffer
	p := New(&buf)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		p.Run(ctx, []config.Hook{}, results)
		wg.Done()
	}()
	cancel()
	wg.Wait()

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

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		p.Run(ctx, hooks, results)
		wg.Done()
	}()
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
	time.Sleep(125 * time.Millisecond)
	cancel()
	wg.Wait()

	require.Contains(t, buf.String(), `| foo | ⣽ |`)
	require.Contains(t, buf.String(), `| bar | ⣽ |`)

	require.Contains(t, buf.String(), `| foo | ✓ |`)
	require.Contains(t, buf.String(), `| bar | ✗ |`)
}
