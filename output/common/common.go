package common

import (
	"context"

	"foreplay/config"
)

type Result struct {
	Hook config.Hook
	Err  error
	Out  []byte
}

type Registerable interface {
	Register(ctx context.Context, hooks []config.Hook, results <-chan Result) chan interface{}
}
