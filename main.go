//go:generate mockery -name=Registerable -outpkg=mockstest -output=mockstest -case=snake -dir=output/common

package main

import (
	"github.com/eberkund/foreplay/cmd"
)

func main() {
	cmd.Execute()
}
