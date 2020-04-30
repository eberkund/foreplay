package spinner

import (
	"bytes"
	"context"
	"testing"

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
	done := p.Register(ctx, []config.Hook{{ID: "foo"}}, results)
	go func() {
		results <- common.Result{
			Hook: config.Hook{
				ID: "foo",
			},
			Err: nil,
			Out: []byte("hello world"),
		}
		cancel()
	}()
	<-done
}
