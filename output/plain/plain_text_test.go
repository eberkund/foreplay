package plain

import (
	"bytes"
	"context"
	"sync"
	"testing"

	"github.com/eberkund/foreplay/config"
	"github.com/eberkund/foreplay/output/common"

	"github.com/stretchr/testify/require"
)

func TestPlainTextOutput(t *testing.T) {
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

func TestPlainPrintsResults(t *testing.T) {
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
	results <- common.Result{
		Hook: config.Hook{},
		Err:  nil,
		Out:  []byte("hello world"),
	}
	cancel()
	wg.Wait()
	require.Contains(t, buf.String(), "hello world")
}
