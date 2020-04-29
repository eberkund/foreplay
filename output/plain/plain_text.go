package plain

import (
	"context"
	"fmt"

	"foreplay/config"
	"foreplay/output/common"
)

type plainPrinter struct{}

func (plainPrinter) Register(ctx context.Context, _ []config.Hook, results <-chan common.Result) chan interface{} {
	done := make(chan interface{})
	go func() {
		defer close(done)
		for {
			select {
			case r := <-results:
				fmt.Println(string(r.Out))
			case <-ctx.Done():
				return
			}
		}
	}()
	return done
}

func New() common.Registerable {
	return &plainPrinter{}
}
