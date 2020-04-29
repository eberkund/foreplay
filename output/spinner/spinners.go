package spinner

import (
	"context"
	"sort"
	"time"

	"foreplay/config"
	"foreplay/output/common"

	"github.com/fatih/color"
	"github.com/k0kubun/go-ansi"
	"github.com/olekukonko/tablewriter"
)

type spinnerPrinter struct {
	spinners map[string]*spinner
	ticker   *time.Ticker
}

type spinner struct {
	success *bool
	ticks   int
}

func New() common.Registerable {
	return &spinnerPrinter{
		spinners: make(map[string]*spinner),
		ticker:   time.NewTicker(125 * time.Millisecond),
	}
}

func (p *spinnerPrinter) createHookJobs(hooks []config.Hook) {
	for _, v := range hooks {
		p.spinners[v.ID] = &spinner{}
	}
}

func (p *spinnerPrinter) Register(ctx context.Context, hooks []config.Hook, results <-chan common.Result) chan interface{} {
	done := make(chan interface{})
	go func() {
		ansi.CursorHide()
		defer ansi.CursorShow()
		defer close(done)

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
	}()
	return done
}

func (p *spinnerPrinter) refresh(reset bool) {
	keys := make([]string, 0)
	for k := range p.spinners {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	table := tablewriter.NewWriter(ansi.NewAnsiStdout())
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
