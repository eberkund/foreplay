//+build wireinject

package run

import (
	"github.com/google/wire"
)

func InitializeRunner() (*Run, error) {
	wire.Build(
		GetConfig,
		GetPrinter,
		GetShell,
		GetRun,
	)
	return &Run{}, nil
}
