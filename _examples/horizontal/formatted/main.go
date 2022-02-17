package main

import (
	"fmt"
	"os"

	"github.com/mergestat/termbar"
)

func main() {
	c := termbar.NewHorizontal(termbar.Bars{
		{Label: "a", Value: 50},
		{Label: "b", Value: 25},
		{Label: "c", Value: 15},
		{Label: "d", Value: 78},
	},
		termbar.WithMaxVal(100),
		termbar.WithLabelSeparator(" - "),
		termbar.WithValueFormatter(func(b termbar.Bar) string { return fmt.Sprintf(" %0.2f%%", b.Value) }),
	)

	c.Print(os.Stdout)
}
