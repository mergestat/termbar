package main

import (
	"os"

	"github.com/mergestat/termbar"
)

func main() {
	c := termbar.NewSparkline(termbar.Bars{
		{Label: "a", Value: 50},
		{Label: "b", Value: 25},
		{Label: "c", Value: 15},
		{Label: "d", Value: 78},
		{Label: "d", Value: 0},
		{Label: "d", Value: 15},
		{Label: "a", Value: 50},
		{Label: "b", Value: 25},
		{Label: "c", Value: 15},
		{Label: "d", Value: 78},
		{Label: "d", Value: 0},
		{Label: "d", Value: 15},
		{Label: "a", Value: 50},
		{Label: "b", Value: 25},
		{Label: "c", Value: 15},
		{Label: "d", Value: 78},
		{Label: "d", Value: 0},
		{Label: "d", Value: 15},
		{Label: "a", Value: 50},
		{Label: "b", Value: 25},
		{Label: "c", Value: 15},
		{Label: "d", Value: 78},
		{Label: "d", Value: 0},
		{Label: "d", Value: 15},
	}, termbar.WithMaxWidth(50), termbar.WithMaxVal(100))

	c.Print(os.Stdout)
}
