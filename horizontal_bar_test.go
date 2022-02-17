package termbar_test

import (
	"os"
	"testing"

	"github.com/mergestat/termbar"
)

func TestHorizontalBar(t *testing.T) {
	c := termbar.NewHorizontal([]termbar.Bar{
		{Label: "a", Value: 50},
		{Label: "b", Value: 25},
		{Label: "b", Value: 15},
	})

	c.Print(os.Stdout)
}
