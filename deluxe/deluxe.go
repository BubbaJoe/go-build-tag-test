//go:build deluxe
// +build deluxe

package deluxe

import (
	"github.com/fatih/color"
)

type Deluxe struct {
}

func NewDeluxe() *Deluxe {
	color.Green("Deluxe is enabled")
	return &Deluxe{}
}
