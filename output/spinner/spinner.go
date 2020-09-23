package spinner

import (
	"github.com/eiannone/keyboard"

	"github.com/fatih/color"
)

type spinner struct {
	success *bool
	ticks   int
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

func discardKeyboardInput() {
	keys, _ := keyboard.GetKeys(1)
	defer keyboard.Close()
	for {
		<-keys
	}
}
