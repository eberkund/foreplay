//go:generate mockery -name=Registerable -outpkg=mockstest -output=mockstest -case=snake -dir=output/common

package main

import (
	"foreplay/cmd"
)

func main() {
	cmd.Execute()
}
