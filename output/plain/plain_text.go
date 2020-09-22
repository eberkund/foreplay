package plain

import (
	"context"
	"fmt"
	"io"

	"github.com/eberkund/foreplay/config"
	"github.com/eberkund/foreplay/output/common"
)

type plainPrinter struct {
	writer io.Writer
}

func (p plainPrinter) Run(ctx context.Context, _ []config.Hook, results <-chan common.Result) {
	for {
		select {
		case r := <-results:
			_, _ = fmt.Fprintln(p.writer, string(r.Out))
		case <-ctx.Done():
			return
		}
	}
}

func New(writer io.Writer) common.Registerable {
	return &plainPrinter{
		writer: writer,
	}
}
