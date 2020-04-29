package plain

import (
	"context"
	"testing"

	"foreplay/config"
	"foreplay/output/common"
)

func TestPlainTextOutput(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	results := make(chan common.Result)
	p := New()
	done := p.Register(ctx, []config.Hook{}, results)
	cancel()
	<-done
}
