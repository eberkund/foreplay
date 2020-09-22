package spinner

import (
	"context"
	"io"
	"sort"
	"time"

	"github.com/eberkund/foreplay/config"
	"github.com/eberkund/foreplay/output/common"

	"github.com/fatih/color"
	"github.com/k0kubun/go-ansi"
	"github.com/olekukonko/tablewriter"
)

type spinnerPrinter struct {
	spinners map[string]*spinner
	ticker   *time.Ticker
	writer   io.Writer
}

type spinner struct {
	success *bool
	ticks   int
}

func New(writer io.Writer) common.Registerable {
	return &spinnerPrinter{
		spinners: make(map[string]*spinner),
		ticker:   time.NewTicker(125 * time.Millisecond),
		writer:   writer,
	}
}

func (p *spinnerPrinter) createHookJobs(hooks []config.Hook) {
	for _, v := range hooks {
		p.spinners[v.ID] = &spinner{}
	}
}

func (p *spinnerPrinter) Run(ctx context.Context, hooks []config.Hook, results <-chan common.Result) {
	ansi.CursorHide()
	defer ansi.CursorShow()
	p.createHookJobs(hooks)

	for {
		select {
		case <-p.ticker.C:
			p.refresh(true)
		case r := <-results:
			success := r.Err == nil
			p.spinners[r.Hook.ID].success = &success
		case <-ctx.Done():
			p.refresh(false)
			return
		}
	}
}

func (p *spinnerPrinter) refresh(reset bool) {
	if len(p.spinners) == 0 {
		return
	}
	keys := make([]string, 0)
	for k := range p.spinners {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	table := tablewriter.NewWriter(p.writer)
	for _, id := range keys {
		spinner := p.spinners[id]
		spinner.ticks++
		table.Append([]string{
			id,
			spinner.progressChar(),
		})
	}
	table.Render()
	if reset {
		ansi.CursorUp(len(p.spinners) + 2)
	}
}

func (h spinner) progressChar() string {
	charSet := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	successSymbol := "✓"
	errorSymbol := "✗"
	if h.success == nil {
		return charSet[h.ticks%len(charSet)]
	}
	if *h.success {
		return color.GreenString(successSymbol)
	}
	return color.RedString(errorSymbol)
}
