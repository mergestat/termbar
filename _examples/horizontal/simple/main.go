package main

import (
	"os"

	"github.com/mergestat/termbar"
)

func main() {
	c := termbar.NewHorizontal(termbar.Bars{
		{Label: "a", Value: 50},
		{Label: "b", Value: 25},
		{Label: "c", Value: 15},
		{Label: "d", Value: 78},
	}, termbar.WithMaxWidth(50), termbar.WithBarChars("▇"))

	c.Print(os.Stdout)
}
