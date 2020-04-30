package plain

import (
	"context"
	"fmt"
	"io"

	"foreplay/config"
	"foreplay/output/common"
)

type plainPrinter struct {
	writer io.Writer
}

func (p plainPrinter) Register(ctx context.Context, _ []config.Hook, results <-chan common.Result) chan interface{} {
	done := make(chan interface{})
	go func() {
		defer close(done)
		for {
			select {
			case r := <-results:
				_, _ = fmt.Fprintln(p.writer, string(r.Out))
			case <-ctx.Done():
				return
			}
		}
	}()
	return done
}

func New(writer io.Writer) common.Registerable {
	return &plainPrinter{
		writer: writer,
	}
}
