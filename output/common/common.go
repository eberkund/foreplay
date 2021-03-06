package common

import (
	"context"

	"github.com/eberkund/foreplay/config"
)

type Result struct {
	Hook config.Hook
	Err  error
	Out  []byte
}

type Registerable interface {
	Run(ctx context.Context, hooks []config.Hook, results <-chan Result)
}
